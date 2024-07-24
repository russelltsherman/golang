package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// simple static file server
// func main() {
// 	dir, _ := os.Getwd()
// 	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(dir))))
// }

// a static file server which accepts parameter flags for port and directory
func main() {
	var dir string
	port := flag.String("port", "8080", "port to serve HTTP on")
	path := flag.String("path", "", "path to serve")
	flag.Parse()

	// if path not provided use working directory as path
	if *path == "" {
		dir, _ = os.Getwd()
	} else {
		dir = *path
	}

	log.Fatal(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(dir))))
}
