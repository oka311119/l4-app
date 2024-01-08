package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
	"github.com/oka311119/l4-app/backend/command/internal/area"
	"github.com/oka311119/l4-app/backend/command/internal/auth"
)

type Handler struct {
    useCase area.UseCase
}

func NewHandler(useCase area.UseCase) *Handler {
    return &Handler {
        useCase: useCase,
    }
}

type createAreaInput struct {
    areaName string `json:"areaName"`
}

func (h *Handler) CreateArea(c *gin.Context) {
    inp := new(createAreaInput)

    if err := c.BindJSON(inp); err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    // Get userID from Context
    userObj, exists := c.Get(auth.CtxUserKey)
    if !exists {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
    }

    // Cast to User entity
    user, ok := userObj.(*entity.User)
    if !ok {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": area.ErrUserIsNotExists.Error()})
        return
    }

    if err := h.useCase.CreateArea(c.Request.Context(), user.ID, inp.areaName); err != nil {
        c.AbortWithStatus(http.StatusInternalServerError)
        return
    }

    c.Status(http.StatusOK)
}

