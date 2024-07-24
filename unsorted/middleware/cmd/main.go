package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang/unsorted/middleware"

	"golang.org/x/net/context"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func panicker(w http.ResponseWriter, r *http.Request) {
	panic(middleware.ErrInvalidID)
}

func withContext(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	bar := ctx.Value("foo")
	w.Write((bar.([]byte)))
}

func main() {
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

	logger := middleware.CreateLogger("section4")

	http.Handle("/", middleware.Time(logger, hello))
	http.Handle("/panic", middleware.Recover(panicker))
	http.Handle("/context", middleware.PassContext(withContext))

	fmt.Println("Starting server at port", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
