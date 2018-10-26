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

func createThrottleKey(user *requestBody) string {
	// Get the current minute
	t := time.Now()
	currentMinute := t.Minute()

	// Create a key
	return user.User + "-" + user.URL + ":" + strconv.Itoa(currentMinute)

}

func throttleHandler(client *redis.Client) func(c echo.Context) error {

	return func(c echo.Context) error {
		fmt.Println("\nRequest!")
		// Capture the request body
		u := new(requestBody)

		// Return a failed status if you can't coerce
		if err := c.Bind(u); err != nil {
			return c.String(http.StatusBadRequest, "Incorrect POST body")
		}

		redisKey := createThrottleKey(u)

		// Increment and set timeout
		pipe := client.TxPipeline()
		incr := pipe.Incr(redisKey)
		pipe.Expire(redisKey, time.Second*59)
		// Execute the pipeline
		_, err := pipe.Exec()
		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal server error")
		}
		currentCount := incr.Val()

		// TODO: acquire thresholds from another service
		if currentCount <= 5 {
			return c.JSON(http.StatusOK, apiResponse{true})
		}
		// Otherwise, you are throttled
		return c.JSON(http.StatusOK, apiResponse{false})

	}
}
