package main

import (
	"context"

	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type CityGeo struct {
}

type CityTitle struct {
	Origin string
	En     string
}

type City struct {
	Id    string
	Title CityTitle
	Geo   CityGeo
}

func main() {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	app := gin.Default()

	router := app.Group("/api")

	router.GET("/city/:country", func(c *gin.Context) {
		country, _ := c.Params.Get("country")
		subString := c.Query("sub")

		cities, err := rdb.Get(ctx, country).Result()

		if err != nil {
			panic(err)
		}

		var city []City

		json.Unmarshal([]byte(cities), &city)

		if subString != "" {
			var filteredCities []City

			for _, v := range city {
				if strings.HasPrefix(v.Title.Origin, subString) || strings.HasPrefix(v.Title.Origin, subString) {
					filteredCities = append(filteredCities, v)
				}
			}

			c.JSON(200, gin.H{
				"cities": filteredCities,
			})

			return
		}

		c.JSON(200, gin.H{
			"cities": city,
		})
	})

	app.Run("0.0.0.0:8080")
}
