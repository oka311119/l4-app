package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oka311119/l4-app/backend/command/internal/handler"
)

func main() {
	r := gin.Default()
	r.POST("/item", handler.CreateItem)
	r.PUT("/item", handler.UpdateItem)
	r.Run() // listen and serve on 0.0.0.0:8080
}
