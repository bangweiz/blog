package user

import (
	"github.com/bangweiz/blog/models"
	"github.com/bangweiz/blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type User struct {
	Api string
}

func (*User) UserRegister(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	password2 := c.PostForm("password2")
	err := util.ValidateUserRegister(username, email, password, password2)
	if len(err) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err,
		})
	} else {
		ok := models.UserRegister(username, email, password)
		if ok {
			c.JSON(http.StatusOK, gin.H{ "message": "ok" })
		} else {
			err["email"] = "Email exists"
			c.JSON(http.StatusBadRequest, gin.H{ "data": err })
		}
	}
}

func (*User) UserLogin(c *gin.Context) {
	email := c.PostForm("email")
	pwd := c.PostForm("password")
	err := util.ValidateUserLogin(email, pwd)
	if len(err) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err,
		})
	} else {
		ok, user := models.FindUser(email, pwd)
		if ok {
			token, errors := util.GenerateToken(user.Username, user.Password, user.Email)
			if errors != nil {
				err["message"] = "Failed to generate a token"
				c.JSON(http.StatusBadRequest, gin.H{ "data": err })
			} else {
				data := make(map[string]interface{})
				data["token"] = token
				data["username"] = user.Username
				data["email"] = user.Email
				if user.Email == "446850579@qq.com" {
					data["isAdmin"] = true
				} else {
					data["isAdmin"] = false
				}
				c.JSON(http.StatusOK, gin.H{ "data": data })
			}
		} else {
			err["message"] = "Email or password is incorrect"
			c.JSON(http.StatusBadRequest, gin.H{ "data": err })
		}
	}
}

func (u *User) ParseToken(c *gin.Context) {
	token := c.Param("token")
	claims, err := util.ParseToken(token)
	if err != nil || time.Now().Unix() > claims.ExpiresAt {
		errMsg := make(map[string]string)
		errMsg["message"] = "Invalid token"
		c.JSON(http.StatusBadRequest, gin.H{ "data": errMsg })
	} else {
		res := make(map[string]interface{})
		res["username"] = claims.Username
		res["email"] = claims.Email
		if claims.Email == "446850579@qq.com" {
			res["isAdmin"] = true
		} else {
			res["isAdmin"] = false
		}
		c.JSON(http.StatusOK, gin.H{ "data": res })
	}
}
