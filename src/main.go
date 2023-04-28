package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album struct:
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// album slice to store initial data
var albums = []album{
	{
		ID:     "1",
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99,
	},
	{
		ID:     "2",
		Title:  "Jeru",
		Artist: "Gerry Mulligan",
		Price:  17.99,
	},
	{
		ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99,
	},
}

// get list of albums:
func getAlbums(ctx *gin.Context) {
	statusCode := http.StatusOK // code 200
	ctx.IndentedJSON(statusCode, albums)
}

// add a new album - from JSON received:
func postAlbums(ctx *gin.Context) {
	var newAlbum album

	err := ctx.BindJSON(&newAlbum)
	if err != nil {
		log.Fatalf("Failed to add a new album!")
		return
	}

	// Append the slice with the new album:
	albums = append(albums, newAlbum)

	// serialize the JSON & add to response: code: 201
	ctx.IndentedJSON(http.StatusCreated, newAlbum)

}

func main() {

	// initialise Gin router:
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	// run the server on port 3000:
	err := router.Run(":3000")
	if err != nil {
		log.Fatalf("[Error] failed to start Gin server due to: " + err.Error())
	}

}
