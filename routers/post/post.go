package post

import (
	"github.com/bangweiz/blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Post struct {
	Api string
}

type post struct {
	AuthorID string `json:"author_id"`
	CategoryID string `json:"category_id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

func (*Post) SavePost(c *gin.Context) {
	var post post
	_ = c.BindJSON(&post)
	models.SavePost(post.Title, post.Content, post.AuthorID, post.CategoryID)
}

func (*Post) FetchPosts(c *gin.Context) {
	data := models.FetchPosts()
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (*Post) FetchPost(c *gin.Context) {
	id := c.Param("id")
	data := models.FetchPost(id)
	c.JSON(http.StatusOK, gin.H{"data": data})
}