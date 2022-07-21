import service from '@/utils/request'

// @Tags UserEntryCard
// @Summary 创建UserEntryCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserEntryCard true "创建UserEntryCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userEntryCard/createUserEntryCard [post]
export const createUserEntryCard = (data) => {
  return service({
    url: '/sold/createUserEntryCard',
    method: 'post',
    data
  })
}

// @Tags UserEntryCard
// @Summary 删除UserEntryCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserEntryCard true "删除UserEntryCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userEntryCard/deleteUserEntryCard [delete]
export const deleteUserEntryCard = (data) => {
  return service({
    url: '/sold/deleteUserEntryCard',
    method: 'delete',
    data
  })
}

// @Tags UserEntryCard
// @Summary 删除UserEntryCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserEntryCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userEntryCard/deleteUserEntryCard [delete]
export const deleteUserEntryCardByIds = (data) => {
  return service({
    url: '/sold/deleteUserEntryCardByIds',
    method: 'delete',
    data
  })
}

// @Tags UserEntryCard
// @Summary 更新UserEntryCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserEntryCard true "更新UserEntryCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userEntryCard/updateUserEntryCard [put]
export const updateUserEntryCard = (data) => {
  return service({
    url: '/sold/updateUserEntryCard',
    method: 'put',
    data
  })
}

// @Tags UserEntryCard
// @Summary 用id查询UserEntryCard
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserEntryCard true "用id查询UserEntryCard"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userEntryCard/findUserEntryCard [get]
export const findUserEntryCard = (params) => {
  return service({
    url: '/sold/findUserEntryCard',
    method: 'get',
    params
  })
}

// @Tags UserEntryCard
// @Summary 分页获取UserEntryCard列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取UserEntryCard列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userEntryCard/getUserEntryCardList [get]
export const getUserEntryCardList = (params) => {
  return service({
    url: '/sold/getUserEntryCardList',
    method: 'get',
    params
  })
}
