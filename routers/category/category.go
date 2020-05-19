package category

import (
	"github.com/bangweiz/blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Category struct {
	Api string
}

func (*Category) AddCategory(c *gin.Context) {
	title := c.PostForm("title")
	err := make(map[string]string, 1)
	if title == "" {
		err["title"] = "Title field is required"
		c.JSON(http.StatusBadRequest, gin.H{ "data": err })
	} else {
		ok, res := models.NewCategory(title)
		if ok {
			c.JSON(http.StatusOK, gin.H{ "data": res })
		} else {
			err["title"] = "Failed to create a new category"
			c.JSON(http.StatusBadRequest, gin.H{ "data": err })
		}
	}
}

func (*Category) FetchCategories(c *gin.Context) {
	data := models.FetchCategories()
	c.JSON(http.StatusOK, gin.H{ "data": data })
}

func (*Category) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	ok := models.DeleteCategory(id)
	if ok {
		data := make(map[string]string, 1)
		data["id"] = id
		c.JSON(http.StatusOK, gin.H{"data": data})
	} else {
		errMsg := make(map[string]string, 1)
		errMsg["message"] = "Failed to delete the category"
		c.JSON(http.StatusBadRequest, gin.H{ "data": errMsg })
	}
}

func (*Category) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ "data": "ok" })
}