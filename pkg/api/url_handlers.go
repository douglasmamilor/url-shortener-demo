package api

import (
	"fmt"
	"hash/maphash"
	"math/rand"
	"net/http"
	"net/url"
	"time"
	"url-shortener/pkg/model"

	"github.com/go-chi/chi/v5"
	"github.com/lucsky/cuid"
)

// URLRoutes sets up the handlers
func (a *API) URLRoutes(router *chi.Mux) http.Handler {
	router.Method("POST", "/shorten", Handler(a.shortenURL))
	router.Method("POST", "/redirect/{shortCode:[a-zA-Z0-9]+}", Handler(a.redirect))

	return router
}

func (a *API) shortenURL(w http.ResponseWriter, r *http.Request) *ServerResponse {
	var shorten model.ShortenURLRequest

	if err := decodeJSONBody(r.Body, &shorten); err != nil {
		return RespondWithError(err, "failed to decode request body", http.StatusInternalServerError)
	}

	if shorten.URL == "" {
		return RespondWithError(nil, "missing url parameter", http.StatusBadRequest)
	}

	if !isURL(shorten.URL) {
		return RespondWithError(nil, "invalid url provided", http.StatusBadRequest)
	}

	shortCode := generateShortCode()

	url := &model.URL{
		ID:          cuid.New(),
		ShortCode:   shortCode,
		OriginalURL: shorten.URL,
		CreatedAt:   time.Now(),
	}

	if err := a.DAL.URLDAL.AddURL(url); err != nil {
		return RespondWithError(err, "could not complete request", http.StatusInternalServerError)
	}

	response := map[string]interface{}{
		"original_url": shorten.URL,
		"short_code":   shortCode,
	}

	return &ServerResponse{Payload: response}
}

func (a *API) redirect(w http.ResponseWriter, r *http.Request) *ServerResponse {
	shortCode := chi.URLParam(r, "shortCode")

	if shortCode == "" {
		return RespondWithError(nil, "missing short_code url parameter", http.StatusBadRequest)
	}

	urlRecord, err := a.DAL.URLDAL.FindByShortCode(shortCode)
	if err != nil {
		return RespondWithError(err, fmt.Sprintf("url not found for code: %v", shortCode), http.StatusNotFound)
	}

	response := map[string]interface{}{
		"original_url": urlRecord.OriginalURL,
		"short_code":   shortCode,
	}

	return &ServerResponse{Payload: response, StatusCode: http.StatusMovedPermanently}
}

func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6

	_ = rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64())))
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
