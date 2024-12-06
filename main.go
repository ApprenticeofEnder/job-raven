package main

import (
	// "fmt"
	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/playwright-community/playwright-go"
)

type company struct {
    ID     string  `json:"id"`
    Name   string  `json:"name"`
}

var companies = []company{
	{ID: "1", Name: "1Password"},
	{ID: "2", Name: "Stripe"},
	{ID: "3", Name: "KOHO"},
	{ID: "4", Name: "Lightspeed"},
}
var nextCompanyId = 5

func main() {
    router := gin.Default()
    router.GET("/companies", listCompanies)
	router.POST("/companies", addCompany)

    router.Run("localhost:8080")
}

func listCompanies(c *gin.Context){
	c.IndentedJSON(http.StatusOK, companies)
}

func addCompany(c *gin.Context) {
	var newCompany company

	if err := c.BindJSON(&newCompany); err != nil {
		return
	}

	companies = append(companies, newCompany)
	c.IndentedJSON(http.StatusCreated, newCompany)
}