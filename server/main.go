package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mdeburnay/waxxed/server/api"
	"errors"
)

type Test struct {
	data []map[string][]struct{}
}

const PORT string = ":3001"

func main() {
	err := godotenv.Load(".env")
	internal.Check(err)

	r := gin.Default()
	r.GET("/", index.home)
	r.Run(`localhost` + PORT)
}
