package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UserAdmLogSearch struct{
    user.UserAdmLog
    request.PageInfo
}
