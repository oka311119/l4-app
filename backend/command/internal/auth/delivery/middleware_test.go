package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/oka311119/l4-app/backend/command/internal/auth"
	"github.com/oka311119/l4-app/backend/command/internal/auth/usecase"
	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
)

func TestAuthMiddleware(t *testing.T) {
	r := gin.Default()
	uc := new(usecase.AuthUseCaseMock)

	r.POST("/api/endpoint", NewAuthMiddleware(uc), func(c *gin.Context){
		c.Status(http.StatusOK)
	})

	// No Auth Header request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/endpoint", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Empty Auth Header request
	w = httptest.NewRecorder()
	req.Header.Set("Authorization", "")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Bearer Auth Header with no token request
	w = httptest.NewRecorder()
	uc.On("ParseToken", "").Return(&entity.User{}, auth.ErrInvalidAccessToken)
	req.Header.Set("Authorization", "Bearer ")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Valid Auth Header
	w = httptest.NewRecorder()
	uc.On("ParseToken", "token").Return(&entity.User{}, nil)
	req.Header.Set("Authorization", "Bearer token")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}