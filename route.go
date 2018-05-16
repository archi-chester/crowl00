package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//
func createRouter() {
	//	создаем роутер
	router := mux.NewRouter()
	//router.HandleFunc("/add_post", addPostFunc).Methods("GET")
	router.HandleFunc("/add-post", addPostFunc).Methods("POST")
	router.HandleFunc("/posts", postsFunc).Methods("GET")
	router.HandleFunc("/upload", uploadFunc).Methods("POST")
	router.HandleFunc("/news", newsFunc).Methods("GET")
	router.HandleFunc("/test", testFunc).Methods("GET")
	router.HandleFunc("/test/{id}", testFunc).Methods("GET")
	router.HandleFunc("/cloud", cloudFunc).Methods("GET")
	router.HandleFunc("/cloud/{id}", cloudFunc).Methods("GET")
	router.HandleFunc("/cloud/{id}", cloudUploadFile).Methods("POST")
	router.HandleFunc("/cloud/{id}/{file}", cloudFunc).Methods("GET")
	router.HandleFunc("/cloud/{id}/{file}/{user}", cloudFunc).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+listeningPort, router))
}
