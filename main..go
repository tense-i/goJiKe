package main

import (
	"gojike/webook/internal/repository"
	"gojike/webook/internal/repository/dao"
	"gojike/webook/internal/service"
	"gojike/webook/internal/web"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	server := initWebServer()
	db := initDB()
	u := initUser(db)
	u.RegisterRoutes(server)
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

func initUser(db *gorm.DB) *web.UserHandler {
	ud := dao.NewUserDAO(db)
	repo := repository.NewUsrRepostory(ud)
	svc := service.NewUserService(repo)
	U := web.NewUserHandler(svc)
	return U
}
