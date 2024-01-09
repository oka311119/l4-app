package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/oka311119/l4-app/backend/command/internal/auth"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc auth.UseCase) {
	h := NewHandler(uc)

	authEp := router.Group("/auth")
	{
		authEp.POST("/sign-up", h.SignUp)
		authEp.POST("/sign-in", h.SignIn)
	}
}
