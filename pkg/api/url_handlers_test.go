package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
	"url-shortener/pkg/api"
	"url-shortener/pkg/config"
	"url-shortener/pkg/dal"
	"url-shortener/pkg/dal/dalfakes"
	"url-shortener/pkg/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

var (
	a             *api.API
	serverHandler http.Handler
	urlDAL        *dalfakes.FakeIURLDAL

	dbMock sqlmock.Sqlmock
)

func TestMain(m *testing.M) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Print("error setting up mock database", err)
		os.Exit(1)
	}

	dbMock = mock
	urlDAL = &dalfakes.FakeIURLDAL{}
	dal := &dal.DAL{
		DB:     sqlx.NewDb(db, "sqlmock"),
		URLDAL: urlDAL,
	}
	a = &api.API{
		Config: &config.Config{},
		DAL:    dal,
	}

	serverHandler = a.SetupServerHandler()
	os.Exit(m.Run())
}

func TestShortenURL(t *testing.T) {
	shortenURLRequestJSON := []byte(`{ "url": "https://docs.docker.com/engine/reference/builder" }`)
	shortenURLRequest := &model.ShortenURLRequest{}
	_ = json.Unmarshal(shortenURLRequestJSON, &shortenURLRequest)
	data, _ := json.Marshal(shortenURLRequest)

	url := "/url/shorten"
	req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))

	urlDAL.AddURLReturns(nil)

	w := httptest.NewRecorder()
	serverHandler.ServeHTTP(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestInvalidURL(t *testing.T) {
	shortenURLRequestJSON := []byte(`{ "url": "some_invalid_url" }`)
	shortenURLRequest := &model.ShortenURLRequest{}
	_ = json.Unmarshal(shortenURLRequestJSON, &shortenURLRequest)
	data, _ := json.Marshal(shortenURLRequest)

	url := "/url/shorten"
	req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))

	w := httptest.NewRecorder()
	serverHandler.ServeHTTP(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestRedirect(t *testing.T) {
	url := "/url/redirect/testShortCode"
	req := httptest.NewRequest(http.MethodPost, url, nil)

	queryReturn := &model.URL{
		ID:          "testID",
		OriginalURL: "http://testURL",
		ShortCode:   "testShortCode",
		CreatedAt:   time.Now(),
	}

	urlDAL.FindByShortCodeReturns(queryReturn, nil)

	w := httptest.NewRecorder()
	serverHandler.ServeHTTP(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	payload, _ := io.ReadAll(resp.Body)

	expected := []byte(`{"original_url":"http://testURL","short_code":"testShortCode"}`)

	assert.Equal(t, http.StatusMovedPermanently, resp.StatusCode)
	assert.Equal(t, expected, payload)
}

func TestRedirectWithMissingParam(t *testing.T) {
	url := "/url/redirect"
	req := httptest.NewRequest(http.MethodPost, url, nil)

	w := httptest.NewRecorder()
	serverHandler.ServeHTTP(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
