package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var ALBUMS []Album

type Album struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Artist    string    `json:"artist"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateAlbums(c *gin.Context) {
	newAlbum := Album{}
	err := c.BindJSON(&newAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	ALBUMS = append(ALBUMS, newAlbum)
	// c.IndentedJSON(http.StatusOK, newAlbum)
	c.JSON(http.StatusOK, newAlbum)
}

func GetAlbums(c *gin.Context) {
	// albums := []Album{}
	// c.IndentedJSON(http.StatusOK, ALBUMS)
	c.JSON(http.StatusOK, ALBUMS)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	// c.JSON(http.StatusOK, id)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("id is :%v", id)})
}

func main() {
	fmt.Println("Main program starts... Gin-Gonic")
	r := gin.Default()
	r.GET("/albums/:id", GetAlbumById)
	r.GET("/albums", GetAlbums)
	r.POST("/albums", CreateAlbums)
	r.Run("localhost:8080")
}

// func main() {
// 	router := gin.Default()
// 	router.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	router.Run()
// }
