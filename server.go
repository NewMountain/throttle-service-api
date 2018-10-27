package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
)

func mkRt(url string) string {
	return ("/api/throttling-service/v1" + url)
}

func makeServer() *echo.Echo {
	// Connect to Redis
	var conn string
	env := os.Getenv("ENV")
	if env == "PROD" {
		conn = "redis:6379"
	} else if env != "DEV" {
		// Throw a fit
		log.Fatal("Please set ENV variable to DEV or PROD")
	} else {
		conn = "localhost:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:     conn,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	e := echo.New()
	e.POST(mkRt("/throttle"), throttleHandler(client))

	return e
}

func main() {
	e := makeServer()
	fmt.Println(":1323/api/throttling-service/v1")
	e.Logger.Fatal(e.Start(":1323"))
}
