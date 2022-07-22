package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
)

type UserSearch struct {
	user.User
	request.PageInfo
}

type UserCreate struct {
	UserName string `form:"userName"`
	Phone    string `form:"phone"`
}
