package controllers

import (
	"context"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/zihangg/go-url-shortener/db"
	e "github.com/zihangg/go-url-shortener/errors"
	"github.com/zihangg/go-url-shortener/helpers"
)

type request struct {
	URL         string `json:"url"`
	CustomShort string `json:"short"`
}

type response struct {
	URL    string    `json:"url"`
	Expiry time.Time `json:"expiry"`
}

func Shorten(c *gin.Context) {
	body := &request{}

	c.ShouldBindJSON(body)

	// Validate if URL is valid
	if !govalidator.IsURL(body.URL) {
		err := e.ErrInvalidUrl
		c.Error(err)
		return
	}

	// Ensure that URL is not from same domain (i.e. trying to shorten the url again)
	if !helpers.NonInternalDomain(body.URL) {
		err := e.ErrInvalidUrl
		c.Error(err)
		return
	}

	id := body.CustomShort
	if id == "" {
		id = helpers.GenerateEncodedURL(rand.Uint64())
	}

	// Check if ID already exists, and if it does, reject (i.e. custom short url)
	client := db.CreateRedisClient()
	ctx := context.Background()
	defer client.Close()

	// No error, value returned hence short url exists
	_, err := client.Get(ctx, id).Result()
	if err == nil {
		c.Error(e.ErrShortUrlAlreadyExists)
		return
	}

	// If short url can be used, add to redis with 5 minutes TTL
	expiry := 5 * 60 * time.Second
	err = client.Set(ctx, id, body.URL, expiry).Err()
	if err != nil {
		c.Error(e.ErrRedisSetError)
		return
	}

	localTime := time.Now()

	resp := &response{
		URL:    os.Getenv("DOMAIN") + "/" + id,
		Expiry: localTime.Add(expiry),
	}

	c.JSON(http.StatusOK, resp)
}
