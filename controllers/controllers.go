package controllers

import (
	"io"
	"net/http"
)

func Init() {
	root()
}

func root() {
	http.HandleFunc("/", healthCheckHandler)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `{"payload":{"status":"success","version":3}}`)
}
