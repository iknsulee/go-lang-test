package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"ksleemodule/cisco/schema"
	"ksleemodule/cisco/tenant"
	"ksleemodule/ksleeutility"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/mso/api/v1/tenants", tenants)
	router.GET("/mso/api/v1/schemas/list-identity", schemas)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func tenants(c *gin.Context) {

	statusCode, responseString, err := tenant.GetAllTenants()
	if err != nil {
		panic("GetAllTenants")
	}

	if statusCode != 200 {
		panic("panic")
	}

	fmt.Printf("[%d][%s]\n", statusCode, ksleeutility.GetPrettyStringFromJSONString(responseString))

	var tenants tenant.CampTenant
	err = json.Unmarshal([]byte(responseString), &tenants)
	if err != nil {
		return
	}

	ksleeutility.PrintPrettyStruct("tenants", tenants)

	c.IndentedJSON(http.StatusCreated, tenants)

}
func schemas(c *gin.Context) {

	statusCode, responseString, err := schema.GetAllSchemas()
	if err != nil {
		panic("GetAllSchemas")
	}

	if statusCode != 200 {
		panic("panic")
	}

	fmt.Printf("[%d][%s]\n", statusCode, ksleeutility.GetPrettyStringFromJSONString(responseString))

	var campSchemaV = schema.CampSchemaV{}
	err = json.Unmarshal([]byte(responseString), &campSchemaV)
	if err != nil {
		return
	}
	ksleeutility.PrintPrettyStruct("campSchemaV", campSchemaV)

	c.IndentedJSON(http.StatusCreated, campSchemaV)

}
