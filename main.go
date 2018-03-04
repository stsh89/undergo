package main

import (
	"./controllers"
	"net/http"
	"os"
)

func main() {
	controllers.Init()

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}
