package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"url-shortener/pkg/config"
	"url-shortener/pkg/dal"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// API holds all of the necessary configuration/deps
type API struct {
	Server *http.Server
	Config *config.Config
	DAL    *dal.DAL
}

// Serve will start the service
func (a *API) Serve() error {
	a.Server = &http.Server{
		Addr:           fmt.Sprintf(":%d", a.Config.Port),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        a.SetupServerHandler(),
		MaxHeaderBytes: 1024 * 1024,
	}

	logrus.Info("API: running...")

	return a.Server.ListenAndServe()
}

// Shutdown stops the service api
func (a *API) Shutdown() error {
	return a.Server.Shutdown(context.Background())
}

// SetupServerHandler handles the configuration of the main server handler
func (a *API) SetupServerHandler() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	mux.Get("/live", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(`{ live: "ok" }`)); err != nil {
			fmt.Println("Unable to write live response")
		}
	})

	mux.Mount("/url", a.URLRoutes(mux))
	return mux
}

func decodeJSONBody(body io.ReadCloser, target interface{}) error {
	defer body.Close()

	if body == nil {
		return fmt.Errorf("missing request body")
	}

	if err := json.NewDecoder(body).Decode(&target); err != nil {
		return errors.Wrapf(err, "failed to parse request json")
	}

	return nil
}

// Handler wraps our http handlers so we can execute some actions before and after a handler is run
type Handler func(w http.ResponseWriter, r *http.Request) *ServerResponse

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := h(w, r)
	// Assume that if there are no errors but the status code is not set, then it must have been
	// an oversight and should be set to a default of http.StatusOK
	if response.StatusCode == 0 && response.Err == nil {
		response.StatusCode = http.StatusOK
	}

	var responseBytes []byte
	var err error
	var marshalErr error

	// If there is NO error, send a payload and a nil error field.
	// If there IS an error, send the error and a nil payload field.
	if response.Err != nil {
		responseBytes, marshalErr = json.Marshal(
			map[string]interface{}{
				"error": ErrorResponse{
					ErrorMessage: response.Message,
					ErrorCode:    response.StatusCode,
				},
			})
	} else {
		responseBytes, marshalErr = json.Marshal(response.Payload)
	}

	// If the marshal failed, report it
	if marshalErr != nil {
		response = RespondWithError(err, "failed to create json response", http.StatusInternalServerError)

		responseBytes, _ = json.Marshal(
			map[string]interface{}{
				"error": ErrorResponse{
					ErrorMessage: response.Message,
					ErrorCode:    response.StatusCode,
				},
			})
	}

	WriteJSONResponse(w, response.StatusCode, responseBytes)
}
