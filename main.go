package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Get a database handle.
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run(":" + port)
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	albums, error := albums()
	if error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	id, error := addAlbum(newAlbum)
	if error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
		return
	}
	album, error := albumByID(id)
	if error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, album)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	n, error := strconv.ParseInt(id, 10, 64)

	if error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Invalid param type"})
		return
	}

	album, error := albumByID(n)
	if error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}
