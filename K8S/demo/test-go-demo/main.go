package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var version = flag.String("v", "v1", "v1")
func main() {
	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		flag.Parse()
		hostname, _ := os.Hostname()

		c.String(http.StatusOK, "This is version:%s running in pod %s",*version,hostname)
	})


	router.Run(":8080")
}
