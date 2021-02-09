package controllers

import (
	"encoding/json"
	"server/src/db"
	"strings"

	"github.com/gin-gonic/gin"
)

type Geo struct {
}

type Title struct {
	Origin string `json:"origin"`
	En     string `json:"en"`
}

type City struct {
	ID    string `json:"id"`
	Title `json:"title"`
	Geo   `json:"geo"`
}

func CityController(c *gin.Context) {
	country, _ := c.Params.Get("country")
	subString := c.Query("sub")

	cities, err := db.RedisClient.Get(country).Result()

	if err != nil {
		panic(err)
	}

	var redisCities []*City

	json.Unmarshal([]byte(cities), &redisCities)

	if subString != "" {
		var filteredCities []*City

		for _, v := range redisCities {
			if strings.HasPrefix(v.Origin, subString) || strings.HasPrefix(v.En, subString) {
				filteredCities = append(filteredCities, v)
			}
		}

		c.JSON(200, gin.H{
			"cities": filteredCities,
		})

		return
	}

	c.JSON(200, gin.H{
		"cities": redisCities,
	})
}
