package main

import (
	"encoding/xml"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/:name", IndexHandler)

	router.Run(":5000")

}
func IndexHandler(c *gin.Context) {

	// name := c.Params.ByName("name")

	// c.JSON(200, gin.H{
	// 	"message": "hello " + name,
	// })

	c.XML(200, Person{FirstName: "YeongSeok", LastName: "Kim"})
}

type Person struct {
	XMLName   xml.Name `xml:person`
	FirstName string   `xml:firstName,attr`
	LastName  string   `xml:lastName,attr`
}

type Recipe struct {
	Name string `json:"name"`

	Tags []string `json:"tags"`

	Ingredients []string `json:"ingredients"`

	Instructions []string `json:"instructions"`

	PublishedAt time.Time `json:"publishedAt"`
}
