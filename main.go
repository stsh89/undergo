package main

import (
	"github.com/stsh89/undergo/controllers"
	"net/http"
	"os"
)

func main() {
	controllers.Init()

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}
