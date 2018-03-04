package main

import (
	"net/http"
	"./controllers"
	"os"
)

func main() {
	controllers.Init()

	port := os.Getenv("PORT")
	http.ListenAndServe(":" + port, nil)
}
