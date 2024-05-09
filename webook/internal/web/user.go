package web

import (
	"gojike/webook/internal/domain"
	"gojike/webook/internal/service"
	"net/http"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	nicknameRegexPattern = "^[a-zA-Z0-9_\\u4e00-\\u9fa5]{3,8}$"
)

type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwdExp   *regexp.Regexp
	nickNameExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {

	//预编译操作
	emailExp := regexp.MustCompile(emailRegexPattern, 0)
	passwdExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	nicknameExp := regexp.MustCompile(nicknameRegexPattern, regexp.None)
	return &UserHandler{
		svc,
		emailExp,
		passwdExp,
		nicknameExp,
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
		//println("邮箱表达式不对")
		//正则表达式不对
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		//println("邮箱不对")
		ctx.String(http.StatusOK, "你的邮箱格式不对")
		return
	}

	ok, err = u.passwdExp.MatchString(req.Password)
	if err != nil {
		//正则表达式不对
		ctx.String(http.StatusOK, "系统错误")
		println("正则表达式格式不对")
	}
	if !ok {
		//println("密码格式 不对")
		ctx.String(http.StatusOK, "密码格式不对")
		return
	}

	err = u.svc.SignUp(ctx, domain.User{
		Email:  req.Email,
		Passwd: req.Password,
	})
	if err == service.ErrUserDuplicateEmail {
		ctx.String(http.StatusOK, "邮箱冲突")
		return
	}

	if err != nil {
		ctx.String(http.StatusBadRequest, "系统异常")
		return
	}
	//这里是数据库操作
	ctx.String(http.StatusOK, "注册成功")

}

func (u *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email  string `json:"email"`
		Passwd string `json:"passwd"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	user, err := u.svc.Login(ctx, req.Email, req.Passwd)
	if err == service.ErrIncalidUersOrPasswd {
		ctx.String(http.StatusOK, "用户名或密码不对")
		return
	}

	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
	}

	//在这登录成功
	//设置session
	sess := sessions.Default(ctx)
	//我可以随便设置值了
	sess.Set("userId", user.Id)

	//Options 实际上控制的是cookie
	sess.Options(sessions.Options{
		Secure:   true,
		HttpOnly: true,
	})
	err = sess.Save()
	if err != nil {
		return
	}
	ctx.String(http.StatusOK, "登录成功")
}

func (u *UserHandler) Edit(ctx *gin.Context) {
	//type UserInfo struct {
	//	NickName string `json:"nickName"`
	//	//Email    string `json:"email"`
	//	//PhoneNum string `json:"phoneNum"`
	//	Birthday string `json:"birthday"`
	//	Aboutme  string `json:"aboutme"`
	//}
	//var userinfoReq UserInfo
	//if err := ctx.Bind(&userinfoReq); err != nil {
	//	ctx.String(http.StatusOK, "g ")
	//	return
	//}

	// if len(userinfoReq.NickName) < 2 || len(userinfoReq.NickName) > 9 {
	// 	ctx.String(http.StatusOK, "昵称长度应大于三")
	// 	return
	// }
	// ok, err := u.nickNameExp.MatchString(userinfoReq.NickName)
	// if err != nil {
	// 	ctx.String(http.StatusOK, "系统错误")
	// 	return
	// }

	// if !ok {
	// 	ctx.String(http.StatusOK, "昵称格式错误")
	// }

	// var birth []string = strings.Split(userinfoReq.Birthday, "-")
	// var y, m, d int
	// y, _ = strconv.Atoi(birth[0])
	// m, _ = strconv.Atoi(birth[1])
	// d, _ = strconv.Atoi(birth[2])
	// if time.Now().Year()-y < 10 || m > 12 || m < 1 || d > 32 || d < 1 {
	// 	ctx.String(http.StatusOK, "生日格式错误")
	// 	return
	// }

	//  if len(userinfoReq.PhoneNum) != 11 {
	// 	ctx.String(http.StatusOK, "电话号码格式错误")
	// 	return
	// }

	// if len(userinfoReq.Aboutme) > 64 {
	// 	ctx.String(http.StatusOK, "个人简介格式错误")
	// 	return
	// }
	//if err := u.svc.Edit(ctx, domain.UserInfo{
	//	NickName: userinfoReq.NickName,
	//	//Email:    userinfoReq.Email,
	//	//PhoneNum: userinfoReq.PhoneNum,
	//	Birthday: userinfoReq.Birthday,
	//	Aboutme:  userinfoReq.Aboutme,
	//}); err != nil {
	//	ctx.String(http.StatusOK, "svc 系统错误")
	//	println("svc 系统错误")
	//	return
	//}
	//
	//ctx.String(http.StatusOK, "修改成功")

}

func (h *UserHandler) Profile(ctx *gin.Context) {
	ctx.String(http.StatusOK, "这是 profile")
}
