import service from '@/utils/request'

// @Tags CardEntry
// @Summary 创建CardEntry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CardEntry true "创建CardEntry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cardEntry/createCardEntry [post]
export const createCardEntry = (data) => {
  return service({
    url: '/cardEntry/createCardEntry',
    method: 'post',
    data
  })
}

// @Tags CardEntry
// @Summary 删除CardEntry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CardEntry true "删除CardEntry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cardEntry/deleteCardEntry [delete]
export const deleteCardEntry = (data) => {
  return service({
    url: '/cardEntry/deleteCardEntry',
    method: 'delete',
    data
  })
}

// @Tags CardEntry
// @Summary 删除CardEntry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CardEntry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /cardEntry/deleteCardEntry [delete]
export const deleteCardEntryByIds = (data) => {
  return service({
    url: '/cardEntry/deleteCardEntryByIds',
    method: 'delete',
    data
  })
}

// @Tags CardEntry
// @Summary 更新CardEntry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CardEntry true "更新CardEntry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /cardEntry/updateCardEntry [put]
export const updateCardEntry = (data) => {
  return service({
    url: '/cardEntry/updateCardEntry',
    method: 'put',
    data
  })
}

// @Tags CardEntry
// @Summary 用id查询CardEntry
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CardEntry true "用id查询CardEntry"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /cardEntry/findCardEntry [get]
export const findCardEntry = (params) => {
  return service({
    url: '/cardEntry/findCardEntry',
    method: 'get',
    params
  })
}

// @Tags CardEntry
// @Summary 分页获取CardEntry列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CardEntry列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /cardEntry/getCardEntryList [get]
export const getCardEntryList = (params) => {
  return service({
    url: '/cardEntry/getCardEntryList',
    method: 'get',
    params
  })
}
