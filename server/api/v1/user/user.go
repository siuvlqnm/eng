package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type UserApi struct {
}

var userService = service.ServiceGroupApp.UserServiceGroup.UserService

// CreateUser 创建User
// @Tags User
// @Summary 创建User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.User true "创建User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/createUser [post]
func (userApi *UserApi) CreateUser(c *gin.Context) {
	var u user.User
	_ = c.ShouldBindJSON(&u)
	u.UserUuid = uuid.NewV4()
	if err := userService.CreateUser(u); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteUser 删除User
// @Tags User
// @Summary 删除User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.User true "删除User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteUser [delete]
func (userApi *UserApi) DeleteUser(c *gin.Context) {
	var user user.User
	_ = c.ShouldBindJSON(&user)
	if err := userService.DeleteUser(user); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserByIds 批量删除User
// @Tags User
// @Summary 批量删除User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /user/deleteUserByIds [delete]
func (userApi *UserApi) DeleteUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := userService.DeleteUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUser 更新User
// @Tags User
// @Summary 更新User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.User true "更新User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /user/updateUser [put]
func (userApi *UserApi) UpdateUser(c *gin.Context) {
	var user user.User
	_ = c.ShouldBindJSON(&user)
	if err := userService.UpdateUser(user); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUser 用id查询User
// @Tags User
// @Summary 用id查询User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.User true "用id查询User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /user/findUser [get]
func (userApi *UserApi) FindUser(c *gin.Context) {
	var user user.User
	_ = c.ShouldBindQuery(&user)
	if reuser, err := userService.GetUser(user.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuser": reuser}, c)
	}
}

// GetUserList 分页获取User列表
// @Tags User
// @Summary 分页获取User列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.UserSearch true "分页获取User列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [get]
func (userApi *UserApi) GetUserList(c *gin.Context) {
	var pageInfo userReq.UserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := userService.GetUserInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
