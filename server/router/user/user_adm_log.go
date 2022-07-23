package user

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserAdmLogRouter struct {
}

// InitUserAdmLogRouter 初始化 UserAdmLog 路由信息
func (s *UserAdmLogRouter) InitUserAdmLogRouter(Router *gin.RouterGroup) {
	userAdmLogRouter := Router.Group("adm").Use(middleware.OperationRecord())
	userAdmLogRouterWithoutRecord := Router.Group("adm")
	var userAdmLogApi = v1.ApiGroupApp.UserApiGroup.UserAdmLogApi
	{
		userAdmLogRouter.POST("createUserAdmLog", userAdmLogApi.CreateUserAdmLog)             // 新建UserAdmLog
		userAdmLogRouter.DELETE("deleteUserAdmLog", userAdmLogApi.DeleteUserAdmLog)           // 删除UserAdmLog
		userAdmLogRouter.DELETE("deleteUserAdmLogByIds", userAdmLogApi.DeleteUserAdmLogByIds) // 批量删除UserAdmLog
		userAdmLogRouter.PUT("updateUserAdmLog", userAdmLogApi.UpdateUserAdmLog)              // 更新UserAdmLog
	}
	{
		userAdmLogRouterWithoutRecord.GET("findUserAdmLog", userAdmLogApi.FindUserAdmLog)       // 根据ID获取UserAdmLog
		userAdmLogRouterWithoutRecord.GET("getUserAdmLogList", userAdmLogApi.GetUserAdmLogList) // 获取UserAdmLog列表
	}
}
