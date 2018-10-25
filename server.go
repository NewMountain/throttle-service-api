package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
)

type apiResponse struct {
	Allow bool `json:"allow"`
}

type requestBody struct {
	User string `json:"user"`
	URL  string `json:"url"`
}

func mkRt(url string) string {
	return ("/api/throttling-service/v1" + url)
}

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	e := echo.New()
	e.POST(mkRt("/throttle"), func(c echo.Context) error {
		// Capture the request body
		u := new(requestBody)

		// Return a failed status if you can't coerce
		if err := c.Bind(u); err != nil {
			return c.String(http.StatusBadRequest, "Incorrect POST body")
		}

		// Get the current minute
		t := time.Now()
		currentMinute := t.Minute()

		// Create a key
		redisKey := u.User + "-" + u.URL + ":" + strconv.Itoa(currentMinute)

		// Increment and set timeout
		pipe := client.TxPipeline()
		incr := pipe.Incr(redisKey)
		pipe.Expire(redisKey, time.Second*59)
		// Execute the pipeline
		_, err := pipe.Exec()
		if err != nil {
			fmt.Println(err)
		}
		currentCount := incr.Val()

		if currentCount <= 5 {
			payload := apiResponse{true}
			return c.JSON(http.StatusOK, payload)
		} else {
			payload := apiResponse{false}
			return c.JSON(http.StatusOK, payload)
		}

	})

	e.Logger.Fatal(e.Start(":1323"))
}
