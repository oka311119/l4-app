package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/oka311119/l4-app/backend/command/internal/area"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc area.UseCase) {
    h := NewHandler(uc)

    areaEp := router.Group("/area")
    {
        areaEp.POST("/create", h.CreateArea)
    }
}
