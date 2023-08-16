package router

import (
	"github.com/gorilla/mux"
	"github.com/thakkarnetram/modules/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()	

	// defining the routes
	router.HandleFunc("/api/v1/movies",controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/v1/movie/create",controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/v1/movie/{id}",controller.MarkedWatch).Methods("PUT")
	router.HandleFunc("/api/v1/movie/delete/{id}",controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/v1/movies/delete",controller.DeleteMovies).Methods("DELETE")
	
	return router
}