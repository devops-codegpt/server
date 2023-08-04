package response

import (
	"time"
)

type UserInfo struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	ZhName   string `json:"zhName"`
	NickName string `json:"nickName"`
	RoleSort uint   `json:"roleSort"`
}

type UserList struct {
	Id        uint      `json:"id"`
	Username  string    `json:"username"`
	ZhName    string    `json:"zhName"`
	Email     string    `json:"email"`
	Status    *uint     `json:"status"`
	RoleId    uint      `json:"roleId"`
	Creator   string    `json:"creator"`
	CreatedAt time.Time `json:"createdAt"`
}
