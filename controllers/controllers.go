package controllers

import (
	"io"
	"net/http"
)

func Init() {
	root()
}

func root() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"payload\": {\"status\":\"success\", \"version\": 3}}")
	})
}
