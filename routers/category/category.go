package category

import (
	"github.com/bangweiz/blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Category struct {
	Api string
}

type category struct {
	Title string `json:"title"`
}

func (*Category) AddCategory(c *gin.Context) {
	var category category
	_ = c.BindJSON(&category)
	err := make(map[string]string, 1)
	if category.Title == "" {
		err["title"] = "Title field is required"
		c.JSON(http.StatusBadRequest, gin.H{ "data": err })
	} else {
		ok, res := models.NewCategory(category.Title)
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