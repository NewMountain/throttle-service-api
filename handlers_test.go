package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCreateThrottleKey(t *testing.T) {
	// Mock a request
	testReqBody := &requestBody{"derpzor", "http://swagger.io"}
	// Get the time
	currMinStr := strconv.Itoa(time.Now().Minute())
	// Mock the key
	mockKey := testReqBody.User + "-" + testReqBody.URL + ":" + currMinStr

	key := createThrottleKey(testReqBody)

	if key != mockKey {
		t.Errorf("Expected " + key)
	}
}

func TestThrottleHandlerGoodReq(t *testing.T) {
	// Setup
	e := echo.New()
	reqBody := `{"user":"test","url":"http://loser.com"}`
	req := httptest.NewRequest(http.MethodPost, mkRt("/throttle"), strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	goodClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Fire those side effects!
	h := throttleHandler(goodClient)

	// Assertions
	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "{\"allow\":true}", rec.Body.String())
	}

}

func TestThrottleHandlerBadReq(t *testing.T) {
	// Setup
	e := echo.New()

	goodClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// TODO: You can give the server complete junk, it's just
	// sensitive to the lack of a JSON body
	badReqBody := `{"foo":"bar","toodles":7}`
	badReq := httptest.NewRequest(http.MethodPost, mkRt("/throttle"), strings.NewReader(badReqBody))
	badRec := httptest.NewRecorder()
	badC := e.NewContext(badReq, badRec)

	// Fire those side effects!
	h := throttleHandler(goodClient)

	// Assertions
	if assert.NoError(t, h(badC)) {
		assert.Equal(t, http.StatusBadRequest, badRec.Code)
		assert.Equal(t, "Incorrect POST body", badRec.Body.String())
	}

}

func TestThrottleHandlerBadDBConn(t *testing.T) {
	// Setup
	e := echo.New()

	reqBody := `{"user":"test","url":"http://loser.com"}`
	req := httptest.NewRequest(http.MethodPost, mkRt("/throttle"), strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	badClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:9736",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Fire those side effects!
	h := throttleHandler(badClient)

	// Assertions
	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, "Internal server error", rec.Body.String())
	}

}
