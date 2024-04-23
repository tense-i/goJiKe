package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginMiddleWareBuiler struct {
	paths []string
}

func NewLoginMiddleWareBuilder() *LoginMiddleWareBuiler {
	return &LoginMiddleWareBuiler{}
}

func (l *LoginMiddleWareBuiler) IgnorePaths(path string) *LoginMiddleWareBuiler {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginMiddleWareBuiler) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}
		//不需要校验的页面
		// if ctx.Request.URL.Path == "/users/login" || ctx.Request.URL.Path == "/users/signup" {
		// 	return
		// }
		sess := sessions.Default(ctx)
		id := sess.Get("userId")
		if id != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
