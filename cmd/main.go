package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	router := http.NewServeMux()

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		rw.Write([]byte("hello"))
	})

	http.ListenAndServe(fmt.Sprintf(":%s", PORT), router)
}
