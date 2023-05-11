package albums

import (
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/verma-kunal/Go-REST-API/src"
)

// add a new album - from JSON received:
func PostAlbums(ctx *gin.Context) {
	var newAlbum src.Album

	err := ctx.BindJSON(&newAlbum)
	if err != nil {
		log.Fatalf("Failed to add a new album!")
		return
	}

	// Append the slice with the new album:
	src.Albums = append(src.Albums, newAlbum)

	// serialize the JSON & add to response: code: 201
	ctx.IndentedJSON(http.StatusCreated, newAlbum)

}