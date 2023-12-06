package main

import (
	"net/http"

	"assignment3/router"
)

func main() {
	r := router.NewRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
