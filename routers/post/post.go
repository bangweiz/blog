package post

import (
	"github.com/bangweiz/blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Post struct {
	Api string
}

func (*Post) SavePost(c *gin.Context) {
	authorID := c.PostForm("author_id")
	categoryID := c.PostForm("category_id")
	title := c.PostForm("title")
	content := c.PostForm("content")

	models.SavePost(title, content, authorID, categoryID)
}

func (*Post) FetchPosts(c *gin.Context) {
	data := models.FetchPosts()
	c.JSON(http.StatusOK, gin.H{"data": data})
}
