package card

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CardEntryRouter struct {
}

// InitCardEntryRouter 初始化 CardEntry 路由信息
func (s *CardEntryRouter) InitCardEntryRouter(Router *gin.RouterGroup) {
	cardEntryRouter := Router.Group("cardEntry").Use(middleware.OperationRecord())
	cardEntryRouterWithoutRecord := Router.Group("cardEntry")
	var cardEntryApi = v1.ApiGroupApp.CardApiGroup.CardEntryApi
	{
		cardEntryRouter.POST("createCardEntry", cardEntryApi.CreateCardEntry)   // 新建CardEntry
		cardEntryRouter.DELETE("deleteCardEntry", cardEntryApi.DeleteCardEntry) // 删除CardEntry
		cardEntryRouter.DELETE("deleteCardEntryByIds", cardEntryApi.DeleteCardEntryByIds) // 批量删除CardEntry
		cardEntryRouter.PUT("updateCardEntry", cardEntryApi.UpdateCardEntry)    // 更新CardEntry
	}
	{
		cardEntryRouterWithoutRecord.GET("findCardEntry", cardEntryApi.FindCardEntry)        // 根据ID获取CardEntry
		cardEntryRouterWithoutRecord.GET("getCardEntryList", cardEntryApi.GetCardEntryList)  // 获取CardEntry列表
	}
}
