package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func postMethod(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Print(value)
	c.JSON(200, gin.H{
		"text": string(value),
	})
}

func getQuery(c *gin.Context) {
	name := c.Query("name") // if the query is (someName? name=Harish&age=21 )
	age := c.Query("age")   // then it will fetch the name and age field values and assign it into name and 							age variables  like name = harish and age = 21
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	server := gin.Default()
	server.GET("/", func(c *gin.Context) { // Using GET Method
		c.JSON(200, gin.H{
			"Name": "PaveenV",
		})
	}).GET("/home", func(c *gin.Context) { // Using Get Method
		c.Writer.WriteString("Hello")
		c.Writer.WriteHeader(http.StatusOK)

	})

	server.POST("./", postMethod) // Using POST Method

	server.POST("./querys", getQuery) // Fetching Data From Query like (someName? name=Harish&age=21 )

	server.GET("./param/:f/:action", func(c *gin.Context) { // Fetching data from url
		fValue := c.Param("f")
		action := c.Param("action")
		c.JSON(200, gin.H{
			"fectedValue": fValue,
			"actions":     action,
		})
	})
	server.Run()
}
© 2021 GitHub, Inc.