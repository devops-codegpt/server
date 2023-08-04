package response

import "time"

type RoleList struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Keyword   string    `json:"keyword"`
	Sort      uint      `json:"sort"`
	Desc      string    `json:"desc"`
	Status    *uint     `json:"status"`
	Creator   string    `json:"creator"`
	CreatedAt time.Time `json:"createdAt"`
}
