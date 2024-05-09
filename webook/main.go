package main

import (
	"github.com/gin-contrib/sessions/redis"
	"gojike/webook/internal/repository"
	"gojike/webook/internal/repository/dao"
	"gojike/webook/internal/service"
	"gojike/webook/internal/web"
	"gojike/webook/internal/web/middleware"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db := initDB()
	server := initWebServer()
	initUser(db, server)
	server.Run(":8080")
}

func initWebServer() *gin.Engine {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		//允许哪个源IP
		AllowOrigins: []string{"http://localhost:3000"},
		//允许的请求方式
		AllowMethods: []string{"POST", "GET"},
		//允许的头部
		AllowHeaders: []string{"Content-Type", "authorization"},
		//ExposeHeaders:    []string{},
		//允许携带的用户信息
		AllowCredentials: true,
		//允许的处理函数
		AllowOriginFunc: func(origin string) bool {
			if strings.Contains(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	//使用cookie为仓管
	//使用基于内存的
	//store := cookie.NewStore([]byte("yf86g8tyYTYHN7WmddNjVUyiZIL7KINsMyZlapNlo4xCNLglpmV161NwI2c5ce1O"), []byte("YsboDqjnGKDB7sDc2PaApK76YxKEQOEip3JfeGJT23hCLe8I7iEHCiGVf0cG8hfD"))

	store, err := redis.NewStore(16, "tcp", "localhost:6379", "", []byte("yf86g8tyYTYHN7WmddNjVUyiZIL7KINsMyZlapNlo4xCNLglpmV161NwI2c5ce1O"), []byte("YsboDqjnGKDB7sDc2PaApK76YxKEQOEip3JfeGJT23hCLe8I7iEHCiGVf0cG8hfD"))
	if err != nil {
		panic(err)
	}
	server.Use(sessions.Sessions("mysession", store))

	server.Use(middleware.NewLoginMiddleWareBuilder().IgnorePaths("/usrs/signup").IgnorePaths("/user/login").Build())
	return server

}
func initDB() *gorm.DB {

	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initUser(db *gorm.DB, server *gin.Engine) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUsrRepostory(ud)
	svc := service.NewUserService(repo)
	U := web.NewUserHandler(svc)
	U.RegisterRoutes(server)
	return U
}
