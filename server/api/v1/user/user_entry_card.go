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

type UserEntryCardApi struct {
}

var userEntryCardService = service.ServiceGroupApp.UserServiceGroup.UserEntryCardService

// CreateUserEntryCard 创建UserEntryCard
// @Tags UserEntryCard
// @Summary 创建UserEntryCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.UserEntryCard true "创建UserEntryCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userEntryCard/createUserEntryCard [post]
func (userEntryCardApi *UserEntryCardApi) CreateUserEntryCard(c *gin.Context) {
	var uec user.UserEntryCard
	_ = c.ShouldBindJSON(&uec)
	ui, uerr := userService.GetUser(uec.UserID)
	ci, cerr := service.ServiceGroupApp.CardServiceGroup.GetCardEntry(uec.CardID)
	ces, cserr := service.ServiceGroupApp.CardServiceGroup.GetCardEntrySpecs(uec.CardID)
	if uerr != nil && cerr != nil && cserr != nil {
		global.GVA_LOG.Error("用户或卡项信息有误!", zap.Error(uerr))
		response.FailWithMessage("用户或卡项信息有误", c)
		return
	}
	uec.UserName = ui.UserName
	uec.Phone = ui.Phone
	uec.CardName = ci.CardName
	uec.CardType = ci.CardType
	uec.TotalPrice = ces.Price

	if ci.CardType == 2 {
		switch ces.DateUnit {
		case 1:
			uec.InitAmt = ces.ValidPeriod
		case 2:
			uec.InitAmt = ces.ValidPeriod * 30
		case 3:
			uec.InitAmt = ces.ValidPeriod * 365
		}
	} else {
		uec.InitAmt = ces.ValidTime
	}

	if uec.GiftAmt != 0 {
		uec.TotalAmt = uec.GiftAmt + uec.InitAmt
	} else {
		uec.TotalAmt = uec.InitAmt
	}
	uec.ContractNumber = uuid.NewV4()
	uec.SurplusAmt = uec.TotalAmt
	if err := userEntryCardService.CreateUserEntryCard(uec); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	userService.UpdateUserBeforeBuyCard(uec.UserID)
	response.OkWithMessage("创建成功", c)
}

// DeleteUserEntryCard 删除UserEntryCard
// @Tags UserEntryCard
// @Summary 删除UserEntryCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.UserEntryCard true "删除UserEntryCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userEntryCard/deleteUserEntryCard [delete]
func (userEntryCardApi *UserEntryCardApi) DeleteUserEntryCard(c *gin.Context) {
	var userEntryCard user.UserEntryCard
	_ = c.ShouldBindJSON(&userEntryCard)
	if err := userEntryCardService.DeleteUserEntryCard(userEntryCard); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserEntryCardByIds 批量删除UserEntryCard
// @Tags UserEntryCard
// @Summary 批量删除UserEntryCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserEntryCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userEntryCard/deleteUserEntryCardByIds [delete]
func (userEntryCardApi *UserEntryCardApi) DeleteUserEntryCardByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := userEntryCardService.DeleteUserEntryCardByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserEntryCard 更新UserEntryCard
// @Tags UserEntryCard
// @Summary 更新UserEntryCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body user.UserEntryCard true "更新UserEntryCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userEntryCard/updateUserEntryCard [put]
func (userEntryCardApi *UserEntryCardApi) UpdateUserEntryCard(c *gin.Context) {
	var userEntryCard user.UserEntryCard
	_ = c.ShouldBindJSON(&userEntryCard)
	if err := userEntryCardService.UpdateUserEntryCard(userEntryCard); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserEntryCard 用id查询UserEntryCard
// @Tags UserEntryCard
// @Summary 用id查询UserEntryCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query user.UserEntryCard true "用id查询UserEntryCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userEntryCard/findUserEntryCard [get]
func (userEntryCardApi *UserEntryCardApi) FindUserEntryCard(c *gin.Context) {
	var userEntryCard user.UserEntryCard
	_ = c.ShouldBindQuery(&userEntryCard)
	if reuserEntryCard, err := userEntryCardService.GetUserEntryCard(userEntryCard.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserEntryCard": reuserEntryCard}, c)
	}
}

// GetUserEntryCardList 分页获取UserEntryCard列表
// @Tags UserEntryCard
// @Summary 分页获取UserEntryCard列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userReq.UserEntryCardSearch true "分页获取UserEntryCard列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userEntryCard/getUserEntryCardList [get]
func (userEntryCardApi *UserEntryCardApi) GetUserEntryCardList(c *gin.Context) {
	var pageInfo userReq.UserEntryCardSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := userEntryCardService.GetUserEntryCardInfoList(pageInfo); err != nil {
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
