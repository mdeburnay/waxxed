package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/waxxed/server/api"
)

func home(c *gin.Context) {
	return "Hello world"
}
