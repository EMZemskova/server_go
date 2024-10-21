package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album struct
type album struct {
	ID     string  `json:id`
	Title  string  `json:title`
	Artist string  `json:artist`
	Price  float64 `json:price`
}

// start data
var albums = []album{
	{ID: "1", Title: "Heartbreak Weather", Artist: "Niall Horan", Price: 58.90},
	{ID: "2", Title: "Fine Line", Artist: "Harry Styles", Price: 73.25},
	{ID: "3", Title: "Good Girl Gone Bad", Artist: "Rihanna", Price: 55.30},
}

// MAIN
func main() {
	//create a router
	router := gin.Default()

	//add an end point GET/albums
	router.GET("/albums", getAlbums)
	//add an endPoint POST/albums
	router.POST("/albums", postAlbums)
	//add an endPoint GET/albums/[id]
	router.GET("/albums/:id", getAlbumById)

	//start server at localhost, port 8080
	router.Run("localhost:8080")
}

// GET all albums (endPoint: Get/albums)
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// POST an album (endPoint: Post/albums)
func postAlbums(c *gin.Context) {
	var newAlbum album
	//add info to newAlbum / find the error
	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}
	//add new album to the slice
	albums = append(albums, newAlbum)
	//send responce with 201 and newAlbum
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// GET an album (endPoint: Get/albums/[id])
func getAlbumById(c *gin.Context) {
	//get id
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			//return 200 + album
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	//return 404 + error message
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
