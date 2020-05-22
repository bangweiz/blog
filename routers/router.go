package routers

import (
	"github.com/bangweiz/blog/middleware"
	"github.com/bangweiz/blog/pkg"
	"github.com/bangweiz/blog/routers/category"
	"github.com/bangweiz/blog/routers/post"
	"github.com/bangweiz/blog/routers/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}
	config.AddAllowHeaders("Authrization")
	r.Use(cors.New(config))

	gin.SetMode(pkg.RunMode)

	userApi := user.User{ Api: "api/user" }
	categoryApi := category.Category{ Api: "api/category" }
	postApi := post.Post{ Api: "api/post" }

	userRouter := r.Group(userApi.Api)
	{
		userRouter.POST("register", userApi.UserRegister)
		userRouter.POST("login", userApi.UserLogin)
		userRouter.GET("parse/:token", userApi.ParseToken)
	}

	categoryRouter := r.Group(categoryApi.Api)
	{
		categoryRouter.GET("all", categoryApi.FetchCategories)
		categoryRouter.GET("test", categoryApi.Test)
	}
	categoryRouter.Use(middleware.JWT())
	{
		categoryRouter.POST("new", categoryApi.AddCategory)
		categoryRouter.DELETE(":id", categoryApi.DeleteCategory)
	}

	postRouter := r.Group(postApi.Api)
	{
		postRouter.GET("all", postApi.FetchPosts)
		postRouter.GET("/detail/:id", postApi.FetchPost)
	}
	postRouter.Use(middleware.JWT())
	{
		postRouter.POST("new", postApi.SavePost)
	}

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return r
}