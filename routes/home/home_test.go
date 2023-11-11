package home

import (
	"golang-company-api/utils"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
)

func TestHomepageHandler(t *testing.T) {
	mockupResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`
	r := utils.SetUpRouter()
	r.GET("/", HomepageHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockupResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}
