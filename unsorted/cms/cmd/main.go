package main

import (
	"fmt"
	"golang/unsorted/cms"
	"log"
	"net/http"
	"os"
)

func main() {
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

	http.HandleFunc("/", cms.ServeIndex)
	http.HandleFunc("/new", cms.HandleNew)
	http.HandleFunc("/page", cms.ServePage)
	http.HandleFunc("/post", cms.ServePost)

	fmt.Println("Starting server at port", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
