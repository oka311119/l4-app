package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oka311119/l4-app/backend/command/internal/api"
)

func main() {
	r := gin.Default()
	r.POST("/item", api.CreateItem)
	r.PUT("/item", api.UpdateItem)
	r.Run() // listen and serve on 0.0.0.0:8080
}
