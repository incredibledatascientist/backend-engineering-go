package domain

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Album struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Artist    string    `json:"artist"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetAlbums(c *gin.Context) []Album {
	albums := []Album{}
	db.Find(&albums)
	return albums
}
