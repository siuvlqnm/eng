package user

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserEntryCardRouter struct {
}

// InitUserEntryCardRouter 初始化 UserEntryCard 路由信息
func (s *UserEntryCardRouter) InitUserEntryCardRouter(Router *gin.RouterGroup) {
	userEntryCardRouter := Router.Group("sold").Use(middleware.OperationRecord())
	userEntryCardRouterWithoutRecord := Router.Group("sold")
	var userEntryCardApi = v1.ApiGroupApp.UserApiGroup.UserEntryCardApi
	{
		userEntryCardRouter.POST("createUserEntryCard", userEntryCardApi.CreateUserEntryCard)             // 新建UserEntryCard
		userEntryCardRouter.DELETE("deleteUserEntryCard", userEntryCardApi.DeleteUserEntryCard)           // 删除UserEntryCard
		userEntryCardRouter.DELETE("deleteUserEntryCardByIds", userEntryCardApi.DeleteUserEntryCardByIds) // 批量删除UserEntryCard
		userEntryCardRouter.PUT("updateUserEntryCard", userEntryCardApi.UpdateUserEntryCard)              // 更新UserEntryCard
	}
	{
		userEntryCardRouterWithoutRecord.GET("findUserEntryCard", userEntryCardApi.FindUserEntryCard)       // 根据ID获取UserEntryCard
		userEntryCardRouterWithoutRecord.GET("getUserEntryCardList", userEntryCardApi.GetUserEntryCardList) // 获取UserEntryCard列表
	}
}
