package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

	// Migrate the schema
	db.AutoMigrate(&Album{})

	router := gin.Default()
	router.GET("/albums", getAlbums(db))
	router.GET("/albums/:id", getAlbumByID(db))
	router.POST("/albums", postAlbums(db))

	router.Run(":" + port)
}

func getAlbums(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var albums []Album
		db.Find(&albums)
		c.IndentedJSON(http.StatusOK, albums)
	}

	return gin.HandlerFunc(fn)

}

func postAlbums(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var newAlbum Album

		// Call BindJSON to bind the received JSON to
		// newAlbum.
		if err := c.BindJSON(&newAlbum); err != nil {
			return
		}

		db.Create(&newAlbum)
		c.IndentedJSON(http.StatusCreated, newAlbum)
	}

	return gin.HandlerFunc(fn)
}

func getAlbumByID(db *gorm.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")
		n, error := strconv.ParseInt(id, 10, 64)

		if error != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Invalid param type"})
			return
		}

		var album Album
		db.First(&album, n)
		if album.ID == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not found."})
			return
		}

		c.IndentedJSON(http.StatusOK, album)
	}

	return gin.HandlerFunc(fn)
}
