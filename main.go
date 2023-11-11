package main

import (
	"golang-company-api/routes/companies"
	"golang-company-api/routes/home"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	home.Router(r)
	companies.Router(r)

	r.Run(":5050")
}
