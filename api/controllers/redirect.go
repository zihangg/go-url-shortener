package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/zihangg/go-url-shortener/db"
	e "github.com/zihangg/go-url-shortener/errors"
)

func Redirect(c *gin.Context) {
	url := c.Param("url")

	client := db.CreateRedisClient()
	ctx := context.Background()
	defer client.Close()

	val, err := client.Get(ctx, url).Result()
	if err == redis.Nil {
		c.Error(e.ErrRedisEmptyResult)
		return
	} else if err != nil {
		c.Error(e.ErrRedisGetError)
	}

	c.Redirect(http.StatusMovedPermanently, val)
}

