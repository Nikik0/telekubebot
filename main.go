package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	f := fnsel(1)
	fmt.Println(f(1, 3))
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func fnsel(n int) func(int, int) int {
	if n == 1 {
		return func(i int, i2 int) int { return i * i2 }
	} else {
		return func(i int, i2 int) int {
			return i + i2
		}
	}
}
