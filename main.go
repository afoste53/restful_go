package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "john Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and CLifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}


func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}


// getAlbums responds w list of all albums as json
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the req body
func postAlbums(c *gin.Context) {
	var newAlbum album

	// call BindJSON to bind received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// add the new album to slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
