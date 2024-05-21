package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pulse227/server-recruit-challenge-sample/api/middleware"
	"github.com/pulse227/server-recruit-challenge-sample/controller"
	"github.com/pulse227/server-recruit-challenge-sample/infra/mysql"
	"github.com/pulse227/server-recruit-challenge-sample/service"
)

func NewRouter(db *sql.DB) *mux.Router {
	singerRepo := mysql.NewSingerRepository(db)
	singerService := service.NewSingerService(singerRepo)
	singerController := controller.NewSingerController(singerService)

	r := mux.NewRouter()

	r.HandleFunc("/singers", singerController.GetSingerListHandler).Methods(http.MethodGet)
	r.HandleFunc("/singers/{id:[0-9]+}", singerController.GetSingerDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/singers", singerController.PostSingerHandler).Methods(http.MethodPost)
	r.HandleFunc("/singers/{id:[0-9]+}", singerController.DeleteSingerHandler).Methods(http.MethodDelete)

	r.Use(middleware.LoggingMiddleware)

	return r
}
