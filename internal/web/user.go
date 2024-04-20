package web

import (
	"gojike/webook/internal/domain"
	"gojike/webook/internal/service"
	"net/http"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc       *service.UserService
	emailExp  *regexp.Regexp
	passwdExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailRegexPattern = "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
		// 和上面比起来，用 ` 看起来就比较清爽
		passwordRegexPattern = "^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[!@\\#$%^&*()_+])[A-Za-z\\d!@\\#$%^&*()_+]{6,}$"
	)
	//预编译操作
	emailExp := regexp.MustCompile(emailRegexPattern, 0)
	passwdExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		&service.UserService{},
		emailExp,
		passwdExp,
	}
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	// REST 风格
	//server.POST("/user", h.SignUp)
	//server.PUT("/user", h.SignUp)
	//server.GET("/users/:username", h.Profile)
	ug := server.Group("/users")
	// POST /users/signup
	ug.POST("/signup", h.SignUp)
	// POST /users/login
	ug.POST("/login", h.Login)
	// POST /users/edit
	ug.POST("/edit", h.Edit)
	// GET /users/profile
	ug.GET("/profile", h.Profile)
}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var req SignUpReq

	//Bind 方法会根据Context-Type来解析你的数据到req里面、解析错误返回400错误
	if err := ctx.Bind(&req); err != nil {
		return
	}

	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		//正则表达式不对
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusBadRequest, "你的邮箱格式不对")
		return
	}

	ok, err = u.passwdExp.MatchString(req.Password)
	if err != nil {
		//正则表达式不对
		ctx.String(http.StatusInternalServerError, "系统错误")
	}
	if !ok {
		ctx.String(http.StatusBadRequest, "密码格式不对")
		return
	}

	err = u.svc.SignUp(ctx, domain.User{
		Email:  req.Email,
		Passwd: req.Password,
	})

	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	//这里是数据库操作
	ctx.String(http.StatusOK, "注册成功")

}

func (h *UserHandler) Login(ctx *gin.Context) {

}

func (h *UserHandler) Edit(ctx *gin.Context) {

}

func (h *UserHandler) Profile(ctx *gin.Context) {
	ctx.String(http.StatusOK, "这是 profile")
}
