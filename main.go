package main

import (
	"net/http"
	"io"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"payload\": {\"status\":\"success\", \"version\": 1}}")
	})

	http.ListenAndServe(":" + port, nil)
}
