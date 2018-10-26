package main

import (
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
)

func mkRt(url string) string {
	return ("/api/throttling-service/v1" + url)
}

func makeServer() *echo.Echo {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	e := echo.New()
	e.POST(mkRt("/throttle"), throttleHandler(client))

	return e
}

func main() {
	e := makeServer()
	e.Logger.Fatal(e.Start(":1323"))
}
