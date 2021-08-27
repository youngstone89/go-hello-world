// Package main Recipes API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /v1
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: John Doe<john.doe@example.com> http://john.doe.com
//
//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json
//     - application/xml
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: KEY
//          in: header
//     oauth2:
//         type: oauth2
//         authorizationUrl: /oauth2/auth
//         tokenUrl: /oauth2/token
//         in: header
//         scopes:
//           bar: foo
//         flow: accessCode
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta
package main

import (
	"context"
	"encoding/xml"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var recipes []Recipe

var ctx context.Context

var err error

var client *mongo.Client

func init() {

	// recipes = make([]Recipe, 0)

	// file, _ := ioutil.ReadFile("recipes.json")

	// _ = json.Unmarshal([]byte(file), &recipes)

	// MONGO_URI="mongodb://admin:password@localhost:27017/test?authSource=admin" MONGO_DATABASE=demo go run main.go
	ctx = context.Background()

	client, err = mongo.Connect(ctx,

		options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err = client.Ping(context.TODO(),

		readpref.Primary()); err != nil {

		log.Fatal(err)

	}

	log.Println("Connected to MongoDB")

	// var listOfRecipes []interface{}

	// for _, recipe := range recipes {

	// 	listOfRecipes = append(listOfRecipes, recipe)

	// }

	// collection := client.Database(os.Getenv(

	// 	"MONGO_DATABASE")).Collection("recipes")

	// insertManyResult, err := collection.InsertMany(

	// 	ctx, listOfRecipes)

	// if err != nil {

	// 	log.Fatal(err)

	// }

	log.Println("Inserted recipes: ",

		len(insertManyResult.InsertedIDs))

}

func main() {

	router := gin.Default()

	router.GET("/:name", IndexHandler)

	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipeHandler)
	router.PUT("/recipes/:id", UpdateRecipeHandler)
	router.DELETE("/recipes/:id", DeleteRecipeHandler)
	router.GET("/recipes/search", SearchRecipesHandler)

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

// swagger:operation GET /recipes recipes listRecipes
// Returns list of recipes
// ---
// produces:
// - application/json
// responses:
//     '200':
//         description: Successful operation
func ListRecipeHandler(c *gin.Context) {

	c.JSON(http.StatusOK, recipes)

}

// swagger:operation PUT /recipes/{id} recipes updateRecipe
// Update an existing recipe
// ---
// parameters:
// - name: id
//   in: path
//   description: ID of the recipe
//   required: true
//   type: string
// produces:
// - application/json
// responses:
//     '200':
//         description: Successful operation
//     '400':
//         description: Invalid input
//     '404':
//         description: Invalid recipe ID
func UpdateRecipeHandler(c *gin.Context) {

	id := c.Param("id")
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	index := -1

	for i := 0; i < len(recipes); i++ {

		if recipes[i].ID == id {

			index = i

		}

	}
	recipe.ID = id
	recipes[index] = recipe

	c.JSON(http.StatusOK, recipe)

}

func DeleteRecipeHandler(c *gin.Context) {

	id := c.Param("id")

	index := -1

	for i := 0; i < len(recipes); i++ {

		if recipes[i].ID == id {

			index = i

		}

	}

	if index == -1 {

		c.JSON(http.StatusNotFound, gin.H{

			"error": "Recipe not found"})

		return

	}
	recipes = append(recipes[:index], recipes[index+1:]...)

	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe has been deleted"})

}
func SearchRecipesHandler(c *gin.Context) {

	tag := c.Query("tag")

	listOfRecipes := make([]Recipe, 0)

	for i := 0; i < len(recipes); i++ {

		found := false

		for _, t := range recipes[i].Tags {

			if strings.EqualFold(t, tag) {

				found = true

			}

		}

		if found {

			listOfRecipes = append(listOfRecipes,

				recipes[i])

		}

	}

	c.JSON(http.StatusOK, listOfRecipes)

}
