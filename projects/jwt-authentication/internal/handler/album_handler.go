package handler

import (
	"bookstore/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type AlbumHandler struct {
// 	store storage.BookStorage
// }

// func NewAlbumHandler(store storage.BookStorage) *AlbumHandler {
// 	return &AlbumHandler{store: store}
// }

// func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
// 	books, err := h.store.GetBooks()
// 	if err != nil {
// 		utils.Error(w, http.StatusInternalServerError, err.Error(), nil)
// 		return
// 	}
// 	utils.Success(w, http.StatusOK, "Albums fetched successfully", books)
// }

func GetAlbums(c *gin.Context) {
	albums := []domain.Album{}
	c.IndentedJSON(http.StatusOK, albums)
}
