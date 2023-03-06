package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"todo-api-go/models"

	"github.com/rs/cors"
)

type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn string
	}
	Logfile string
}

type Application struct {
	Config Config
	Logger *log.Logger
	Models models.Models
}

type JsonResp struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func ServerSetting(cfg Config, logger *log.Logger, db *sql.DB) (srv *http.Server) {
	app := &Application{
		Config: cfg,
		Logger: logger,
		Models: models.NewModels(db),
	}
	handler := cors.Default().Handler(app.routes())

	srv = &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      handler,
		IdleTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return
}

func (app *Application) WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	// wrapper := make(map[string]interface{})

	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *Application) ErrorJSON(w http.ResponseWriter, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	app.WriteJSON(w, http.StatusBadRequest, theError)
}
