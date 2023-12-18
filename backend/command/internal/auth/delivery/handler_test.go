package delivery

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignUp(t *testing.T) {
	r := gin.Default()
	uc := new(usercase.AuthUseCaseMock)

	RegisterHTTPEndpoints(r, uc)

	signUpBody := &signInput{
		Username: "testuser",
		Password: "testpass",
	}

	body, err := json.Marshal(signUpBody)
	assert.NoError(t, err)

	uc.On("SignUp", signUpBody.Username, signUpBody.Password).Return(nil)

	w := httptest.NewRecoder()
	req, _ = http.NewRequest("POST", "/auth/sign-up", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestSignIn(t *testing.T) {
	r := gin.Default()
	uc := new(usecase.AuthUseCaseMock)

	RegisterHTTPEndpoints(r, uc)

	signInBody := &signInput{
		Username: "testuser",
		Password: "testpass",
	}

	body, err := json.Marshal(signUpBody)
	assert.NoError(t, err)

	uc.On("SignIn", signUpBody.Username, signInBody.Password).Return("jwt", nil)

	w := httptest.NewRecoder()
	req, _ := http.NewRequest("POST", "/auth/sign-in", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"token\":\"jwt\"}", w.Body.string())
}