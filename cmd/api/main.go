package main

import (
	"encoding/json"
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

	fs := http.FileServer(http.Dir("./cmd/web/public"))

	router.HandleFunc("/api/test", func(rw http.ResponseWriter, r *http.Request) {

		rw.Header().Set("Content-Type", "application/json")

		json.NewEncoder(rw).Encode(map[string]interface{}{
			"status":  200,
			"message": "test",
		})

	})

	router.Handle("/", fs)

	http.ListenAndServe(fmt.Sprintf(":%s", PORT), router)
}
