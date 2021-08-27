package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)

	file, _ := ioutil.ReadFile("recipes.json")

	_ = json.Unmarshal([]byte(file), &recipes)
}

func main() {

	router := gin.Default()

	router.GET("/:name", IndexHandler)

	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipeHandler)

	router.Run()

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
	ID string `json:"id"`

	Name string `json:"name"`

	Tags []string `json:"tags"`

	Ingredients []string `json:"ingredients"`

	Instructions []string `json:"instructions"`

	PublishedAt time.Time `json:"publishedAt"`
}

func NewRecipeHandler(c *gin.Context) {

	var recipe Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})

		return
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)

}

func ListRecipeHandler(c *gin.Context) {

	c.JSON(http.StatusOK, recipes)

}
