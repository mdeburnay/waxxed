package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/waxxed/api"
)

func Home(c *gin.Context) {
	err := godotenv.Load(".env")
	api.Check(err)

	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
}
