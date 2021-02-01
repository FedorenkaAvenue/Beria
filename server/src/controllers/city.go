package controllers

import (
	"fmt"
	"net/http"
)

// City sucks too
type City struct{}

// CityController sucks
func CityController() *City {
	return &City{}
}

func (c *City) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, cities")
}
