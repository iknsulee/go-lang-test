package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

var users = []User{
	{Name: "test", Age: 11},
}

func main() {
	r := setupRouter()
	r.Run(":8080")

}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/someGet", someMethod)
	r.POST("/somePost", someMethod)
	r.PUT("/somePut", someMethod)
	r.DELETE("/someDelete", someMethod)
	r.PATCH("/somePatch", someMethod)
	r.HEAD("/someHead", someMethod)
	r.OPTIONS("/someOptions", someMethod)

	v1 := r.Group("camp/v1")
	v1.POST("/users", PostUser)

	return r
}

func PostUser(c *gin.Context) {

	var newUser User

	//err := c.ShouldBindJSON(newUser)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)

}

func someMethod(context *gin.Context) {
	httpMethod := context.Request.Method
	context.JSON(200, gin.H{"status": "good", "sending": httpMethod})
}
