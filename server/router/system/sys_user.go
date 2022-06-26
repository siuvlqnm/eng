package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysUserRouter struct{}

func (s *SysUserRouter) InitSysUserRouter(Router *gin.RouterGroup) {
	sysUserRouter := Router.Group("staff").Use(middleware.OperationRecord())
	sysyUserRouterWithoutRecord := Router.Group("staff")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		sysUserRouter.POST("admin_register", baseApi.Register)               // 管理员注册账号
		sysUserRouter.POST("changePassword", baseApi.ChangePassword)         // 用户修改密码
		sysUserRouter.POST("setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
		sysUserRouter.DELETE("deleteUser", baseApi.DeleteUser)               // 删除用户
		sysUserRouter.PUT("setUserInfo", baseApi.SetUserInfo)                // 设置用户信息
		sysUserRouter.PUT("setSelfInfo", baseApi.SetSelfInfo)                // 设置自身信息
		sysUserRouter.POST("setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
		sysUserRouter.POST("resetPassword", baseApi.ResetPassword)           // 设置用户权限组
	}
	{
		sysyUserRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // 分页获取用户列表
		sysyUserRouterWithoutRecord.GET("getUserInfo", baseApi.GetUserInfo)  // 获取自身信息
	}
}
