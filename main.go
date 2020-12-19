package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}

func PostHomePage(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "Post Home Page",
	})
}
func QueryStrings(c *gin.Context)  {
	name := c.Query("name") //name=nextwebb
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}

func PathParamenters(c *gin.Context)  {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age": age,
	})
}
func main() {
	fmt.Println("Hello world")

	// initialized the gin router
	r := gin.Default()
	// path,Invoke
	r.GET("/", HomePage )
	r.POST("/", PostHomePage )
	//r.OPTIONS("/", PostHomePage ) // other verbs
	r.GET("/query", QueryStrings ) //query?name=nextwebb&age=35
	r.GET("/path/:name/:age", PathParamenters) // path/nextwebb/35
	r.Run()
}