module github.com/bangweiz/blog

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.2
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20200430082407-1f5687305801 // indirect
	gopkg.in/ini.v1 v1.55.0
)

replace (
	github.com/bangweiz/blog/middleware => ./blog/middleware
	github.com/bangweiz/blog/models => ./blog/models
	github.com/bangweiz/blog/pkg => ./blog/pkg
	github.com/bangweiz/blog/routers => ./blog/routers
	github.com/bangweiz/blog/routers/category => ./blog/routers/category
	github.com/bangweiz/blog/routers/post => ./blog/routers/post
	github.com/bangweiz/blog/routers/user => ./blog/routers/user
)
