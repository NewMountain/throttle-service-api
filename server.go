package main

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
)

func mkRt(url string) string {
	return ("/api/throttling-service/v1" + url)
}

func makeServer() *echo.Echo {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "redis",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	e := echo.New()
	e.POST(mkRt("/throttle"), throttleHandler(client))

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return e
}

func main() {
	e := makeServer()
	e.Logger.Fatal(e.Start(":1323"))
}
