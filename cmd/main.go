package main

import "net/http"

func appHandler() http.Handler {
	return
}

func main() {
	server := http.Server{
		Addr:    ":8081",
		Handler: appHandler(),
	}
	server.ListenAndServe()
}
