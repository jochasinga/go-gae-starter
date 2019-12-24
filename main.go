package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jochasinga/boo/model"
	"github.com/jochasinga/boo/route"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/greet" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Greeting!")
}

func init() {
	model.Initialize()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	router := route.NewRouter()
	router.Run(":" + port)
}
