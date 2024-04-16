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
		{0, "hello", 0},
		{0, "hello", 0},
		{0, "hello", 0},
		{0, "hello", 0},
		{0, "hello", 0},
		{0, "hello", 0},
		{0, "hello", 0},
		{0, "hello", 0},
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    p,
	},
	)
}
