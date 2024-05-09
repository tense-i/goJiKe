package middleware

import (
	"encoding/gob"
	"net/http"
	"time"

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
	// 用GO的方式编码解码
	gob.Register(time.Now())
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

		updatetime := sess.Get("update_time")
		//要放回去
		sess.Set("userId", id)
		now := time.Now()

		//说明刚登录、没被刷新过
		if updatetime == nil {
			sess.Set("update_time", now)
			sess.Save()
			return
		}
		//刷新过
		updateTimeVal, _ := updatetime.(time.Time)
		if now.Sub(updateTimeVal) > time.Minute {
			sess.Set("update_time", now)
			sess.Save()
			return
		}
	}
}
