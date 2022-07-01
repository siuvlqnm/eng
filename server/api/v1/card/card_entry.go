package card

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/card"
	cardReq "github.com/flipped-aurora/gin-vue-admin/server/model/card/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CardEntryApi struct {
}

var s = service.ServiceGroupApp.CardServiceGroup.CardEntryService
var ces = service.ServiceGroupApp.CardServiceGroup.CardEntrySpecsService

// CreateCardEntry 创建CardEntry
// @Tags CardEntry
// @Summary 创建CardEntry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body card.CardEntry true "创建CardEntry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cardEntry/createCardEntry [post]
func (i *CardEntryApi) CreateCardEntry(c *gin.Context) {
	var ce cardReq.CardEntryCreate
	_ = c.ShouldBindJSON(&ce)
	if cardId, err := s.CreateCardEntry(ce); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		ce.CardEntryId = cardId
		ces.CreateCardEntrySpecs(ce)
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCardEntry 删除CardEntry
// @Tags CardEntry
// @Summary 删除CardEntry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body card.CardEntry true "删除CardEntry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cardEntry/deleteCardEntry [delete]
func (i *CardEntryApi) DeleteCardEntry(c *gin.Context) {
	var cardEntry card.CardEntry
	_ = c.ShouldBindJSON(&cardEntry)
	if err := s.DeleteCardEntry(cardEntry); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCardEntryByIds 批量删除CardEntry
// @Tags CardEntry
// @Summary 批量删除CardEntry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CardEntry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /cardEntry/deleteCardEntryByIds [delete]
func (i *CardEntryApi) DeleteCardEntryByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := s.DeleteCardEntryByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCardEntry 更新CardEntry
// @Tags CardEntry
// @Summary 更新CardEntry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body card.CardEntry true "更新CardEntry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cardEntry/updateCardEntry [put]
func (i *CardEntryApi) UpdateCardEntry(c *gin.Context) {
	var cardEntry card.CardEntry
	_ = c.ShouldBindJSON(&cardEntry)
	if err := s.UpdateCardEntry(cardEntry); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCardEntry 用id查询CardEntry
// @Tags CardEntry
// @Summary 用id查询CardEntry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query card.CardEntry true "用id查询CardEntry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cardEntry/findCardEntry [get]
func (cardEntryApi *CardEntryApi) FindCardEntry(c *gin.Context) {
	var cardEntry card.CardEntry
	_ = c.ShouldBindQuery(&cardEntry)
	if recardEntry, err := s.GetCardEntry(cardEntry.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recardEntry": recardEntry}, c)
	}
}

// GetCardEntryList 分页获取CardEntry列表
// @Tags CardEntry
// @Summary 分页获取CardEntry列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cardReq.CardEntrySearch true "分页获取CardEntry列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cardEntry/getCardEntryList [get]
func (i *CardEntryApi) GetCardEntryList(c *gin.Context) {
	var pageInfo cardReq.CardEntrySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := s.GetCardEntryInfoList(pageInfo); err != nil {
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
