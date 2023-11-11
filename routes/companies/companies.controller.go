package companies

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Company struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	CEO     string `json:"ceo"`
	Revenue string `json:"revenue"`
}

func Router(r *gin.Engine) {
	companyRoutes := r.Group("/companies")
	companyRoutes.GET("/", GetCompaniesHandler)
	companyRoutes.POST("/", NewCompanyHandler)
	companyRoutes.GET("/:id", GetCompanyHandler)
	companyRoutes.PUT("/:id", UpdateCompanyHandler)
	companyRoutes.DELETE("/:id", DeleteCompanyHandler)
}

func NewCompanyHandler(c *gin.Context) {
	var newCompany Company
	if err := c.ShouldBindJSON(&newCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	New(newCompany)
	c.JSON(http.StatusCreated, newCompany)
}

func GetCompaniesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, companies)
}

func GetCompanyHandler(c *gin.Context) {
	id := c.Param("id")

	if company := FindById(id); company != nil {
		c.JSON(http.StatusOK, company)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Company not found!"})
}

func UpdateCompanyHandler(c *gin.Context) {
	id := c.Param("id")
	var updatedCompany Company
	if err := c.ShouldBindJSON(&updatedCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company, updated := UpdateById(id, updatedCompany)

	if !updated {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, company)
}

func DeleteCompanyHandler(c *gin.Context) {
	id := c.Param("id")
	deleted := Delete(id)

	if !deleted {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found!"})
		return
	}

	c.JSON(http.StatusOK, companies)
}
