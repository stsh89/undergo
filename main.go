package main

import (
	"net/http"
	"io"
)

func main() {
	port := "3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"payload\": {\"status\":\"success\"}}")
	})

	http.ListenAndServe(":" + port, nil)
}
