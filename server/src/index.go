package main

import (
	"controllers"
	"net/http"
)

func main() {
	sm := http.NewServeMux()
	sm.Handle("/", controllers.CityController())

	http.ListenAndServe(":8080", sm)
}
