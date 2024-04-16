package controller

import (
	"gin_test/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToiletList(c *gin.Context) {
	//bookService :=service.BookService{}
	//BookLists := bookService.GetBookList()

	var p = []*model.Toilet{
		{0, "apple", 0},
		{1, "banana", 0},
		{2, "camel", 0},
		{3, "doctor", 0},
		{4, "ero", 0},
		{5, "fuck", 0},
		{6, "g-spot", 0},
		{7, "H", 0},
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    p,
	},
	)
}
