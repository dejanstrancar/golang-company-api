package companies

import (
	"bytes"
	"encoding/json"
	"golang-company-api/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
	"github.com/rs/xid"
)

func TestNewCompanyHandler(t *testing.T) {
	r := utils.SetUpRouter()
	r.POST("/company", NewCompanyHandler)

	companyId := xid.New().String()
	company := Company{
		ID:      companyId,
		Name:    "Demo Company",
		CEO:     "Demo CEO",
		Revenue: "35 million",
	}
	jsonValue, _ := json.Marshal(company)
	req, _ := http.NewRequest("POST", "/company", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

}

func TestGetCompaniesHandler(t *testing.T) {
	r := utils.SetUpRouter()
	r.GET("/companies", GetCompaniesHandler)
	req, _ := http.NewRequest("GET", "/companies", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var companies []Company
	json.Unmarshal(w.Body.Bytes(), &companies)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEqual(t, 0, len(companies))
}

func TestUpdateCompanyHandler(t *testing.T) {
	r := utils.SetUpRouter()
	r.PUT("/company/:id", UpdateCompanyHandler)
	company := Company{
		ID:      `2`,
		Name:    "Demo Company",
		CEO:     "Demo CEO",
		Revenue: "35 million",
	}
	jsonValue, _ := json.Marshal(company)
	reqFound, _ := http.NewRequest("PUT", "/company/"+company.ID, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateCompanyHandlerNotFound(t *testing.T) {
	r := utils.SetUpRouter()
	r.PUT("/company/:id", UpdateCompanyHandler)
	company := Company{
		ID:      `2`,
		Name:    "Demo Company",
		CEO:     "Demo CEO",
		Revenue: "35 million",
	}
	jsonValue, _ := json.Marshal(company)

	reqNotFound, _ := http.NewRequest("PUT", "/company/12", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
