package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping is the router to check if this webservice is working
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
