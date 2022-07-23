package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserAdmLogApi struct {
}

var userAdmLogService = service.ServiceGroupApp.UserServiceGroup.UserAdmLogService

// CreateUserAdmLog 创建UserAdmLog
// @Tags UserAdmLog
// @Summary 创建UserAdmLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.UserAdmLog true "创建UserAdmLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAdmLog/createUserAdmLog [post]
func (userAdmLogApi *UserAdmLogApi) CreateUserAdmLog(c *gin.Context) {
	var ual user.UserAdmLog
	_ = c.ShouldBindJSON(&ual)

	ui, _ := userService.GetUser(ual.UserID)
	ual.UserName = ui.UserName
	ual.Phone = ui.Phone

	ueci, _ := userEntryCardService.GetUserEntryCard(ual.UserCardID)
	ual.CardName = ueci.CardName

	if err := userAdmLogService.CreateUserAdmLog(ual); err != nil {
		global.GVA_LOG.Error("入场失败!", zap.Error(err))
		response.FailWithMessage("入场失败", c)
		return
	}

	userEntryCardService.SuccEntry(ueci)
	response.OkWithMessage("入场成功", c)
}

// DeleteUserAdmLog 删除UserAdmLog
// @Tags UserAdmLog
// @Summary 删除UserAdmLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.UserAdmLog true "删除UserAdmLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAdmLog/deleteUserAdmLog [delete]
func (userAdmLogApi *UserAdmLogApi) DeleteUserAdmLog(c *gin.Context) {
	var userAdmLog user.UserAdmLog
	_ = c.ShouldBindJSON(&userAdmLog)
	if err := userAdmLogService.DeleteUserAdmLog(userAdmLog); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserAdmLogByIds 批量删除UserAdmLog
// @Tags UserAdmLog
// @Summary 批量删除UserAdmLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserAdmLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userAdmLog/deleteUserAdmLogByIds [delete]
func (userAdmLogApi *UserAdmLogApi) DeleteUserAdmLogByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := userAdmLogService.DeleteUserAdmLogByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserAdmLog 更新UserAdmLog
// @Tags UserAdmLog
// @Summary 更新UserAdmLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.UserAdmLog true "更新UserAdmLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userAdmLog/updateUserAdmLog [put]
func (userAdmLogApi *UserAdmLogApi) UpdateUserAdmLog(c *gin.Context) {
	var userAdmLog user.UserAdmLog
	_ = c.ShouldBindJSON(&userAdmLog)
	if err := userAdmLogService.UpdateUserAdmLog(userAdmLog); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserAdmLog 用id查询UserAdmLog
// @Tags UserAdmLog
// @Summary 用id查询UserAdmLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.UserAdmLog true "用id查询UserAdmLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userAdmLog/findUserAdmLog [get]
func (userAdmLogApi *UserAdmLogApi) FindUserAdmLog(c *gin.Context) {
	var userAdmLog user.UserAdmLog
	_ = c.ShouldBindQuery(&userAdmLog)
	if reuserAdmLog, err := userAdmLogService.GetUserAdmLog(userAdmLog.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserAdmLog": reuserAdmLog}, c)
	}
}

// GetUserAdmLogList 分页获取UserAdmLog列表
// @Tags UserAdmLog
// @Summary 分页获取UserAdmLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.UserAdmLogSearch true "分页获取UserAdmLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAdmLog/getUserAdmLogList [get]
func (userAdmLogApi *UserAdmLogApi) GetUserAdmLogList(c *gin.Context) {
	var pageInfo userReq.UserAdmLogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := userAdmLogService.GetUserAdmLogInfoList(pageInfo); err != nil {
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
