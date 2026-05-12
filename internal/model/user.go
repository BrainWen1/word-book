// model/user.go
// 存放用户相关的结构体定义
package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
