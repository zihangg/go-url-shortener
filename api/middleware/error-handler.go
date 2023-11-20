package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	e "github.com/zihangg/go-url-shortener/errors"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
			switch err.Err.Error() {
			case e.ErrInvalidUrl.Error():
					c.JSON(http.StatusBadRequest, gin.H{
							"error": "Invalid URL!",
					})
					return
			case e.ErrShortUrlAlreadyExists.Error():
					c.JSON(http.StatusBadRequest, gin.H{
							"error": "Short URL already exists!",
					})
			default:
					c.JSON(http.StatusInternalServerError, gin.H{
							"error": "Internal Server Error",
					})
					return
			}
	}
}
