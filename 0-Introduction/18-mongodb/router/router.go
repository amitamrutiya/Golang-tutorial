package router

import (
	"mongodb/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", controller.ServeHome).Methods("GET")
	router.HandleFunc("/api/movie", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controller.GetMovie).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movie", controller.DeleteAllMovies).Methods("DELETE")
	router.HandleFunc("/api/movie/watch/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.UpdateMovie).Methods("PUT")

	return router
}
