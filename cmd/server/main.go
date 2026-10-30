package main

import (
	"log"
	"net/http"

	"github.com/parthsolanke/SortViz/internal/handler"
)

func main() {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("web/static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handler.IndexHandler)
	router.HandleFunc("/algorithms", handler.AlgorithmsHandler)
	router.HandleFunc("/about", handler.AboutHandler)
	router.HandleFunc("POST /sort", handler.SortHandler)

	log.Println("Server started at http://127.0.0.1:8080/")
	log.Fatal(http.ListenAndServe(":8080", router))
}
