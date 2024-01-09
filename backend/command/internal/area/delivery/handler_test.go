package delivery

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/oka311119/l4-app/backend/command/internal/domain/entity"
	"github.com/oka311119/l4-app/backend/command/internal/area/usecase"
	"github.com/oka311119/l4-app/backend/command/internal/auth"
)

func TestCreateArea(t *testing.T) {
    testUser := &entity.User{
		Username: "testuser",
		Password: "testpass",
	}

	r := gin.Default()
    group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, testUser)
	})

	uc := new(usecase.AreaUseCaseMock)

	RegisterHTTPEndpoints(group, uc)

	b := &createAreaInput{
        AreaName: "area-name",
	}

	body, err := json.Marshal(b)
	assert.NoError(t, err)

	uc.On("CreateArea", b.AreaName).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/area/create", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

