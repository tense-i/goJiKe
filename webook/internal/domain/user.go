package domain

import "time"

// User 领域对象 是DDD 中得到的 entity
type User struct {
	Id     int64
	Email  string
	Passwd string
	Ctime  time.Time
}

type UserInfo struct {
	NickName string
	//Email    string
	//PhoneNum string
	Birthday string
	Aboutme  string
}
