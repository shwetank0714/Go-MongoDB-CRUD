package router

import (
	"github.com/gorilla/mux"
	controller "github.com/shwetank0714/mongodbapi/controllers"
)


func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies/create", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/mark-watched/{id}", controller.MakrMovieAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/delete-one/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movies/delete-all", controller.DeleteAllMovies).Methods("DELETE")

	return router
}