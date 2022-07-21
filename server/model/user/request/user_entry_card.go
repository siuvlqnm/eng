package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
)

type UserEntryCardSearch struct {
	user.UserEntryCard
	request.PageInfo
}
