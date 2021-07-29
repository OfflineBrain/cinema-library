package main

import (
	"cinema-library/endpoint"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	filmEndpoint := &endpoint.FilmEndpoint{}
	filmEndpoint.Register(router.HandleFunc)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
