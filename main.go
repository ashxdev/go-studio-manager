package main

import (
	"go-dls/handlers"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//cacheControlHeader := "public, max-age=31536000, s-maxage=31536000"

	staticRouter := http.NewServeMux()
	staticFileServer := http.FileServer(http.Dir("./public"))
	staticRouter.Handle("/public/", http.StripPrefix("/public/", staticFileServer))

	router := http.NewServeMux()

	router.HandleFunc("/", handlers.Make(handlers.HandleHomeIndex))
	router.Handle("/public/", staticRouter)

	port := os.Getenv("HTTP_LISTEN_PORT")
	slog.Info("application running", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
