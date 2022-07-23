import service from '@/utils/request'

// @Tags UserAdmLog
// @Summary 创建UserAdmLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAdmLog true "创建UserAdmLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAdmLog/createUserAdmLog [post]
export const createUserAdmLog = (data) => {
  return service({
    url: '/adm/createUserAdmLog',
    method: 'post',
    data
  })
}

// @Tags UserAdmLog
// @Summary 删除UserAdmLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAdmLog true "删除UserAdmLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAdmLog/deleteUserAdmLog [delete]
export const deleteUserAdmLog = (data) => {
  return service({
    url: '/adm/deleteUserAdmLog',
    method: 'delete',
    data
  })
}

// @Tags UserAdmLog
// @Summary 删除UserAdmLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserAdmLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userAdmLog/deleteUserAdmLog [delete]
export const deleteUserAdmLogByIds = (data) => {
  return service({
    url: '/adm/deleteUserAdmLogByIds',
    method: 'delete',
    data
  })
}

// @Tags UserAdmLog
// @Summary 更新UserAdmLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserAdmLog true "更新UserAdmLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userAdmLog/updateUserAdmLog [put]
export const updateUserAdmLog = (data) => {
  return service({
    url: '/adm/updateUserAdmLog',
    method: 'put',
    data
  })
}

// @Tags UserAdmLog
// @Summary 用id查询UserAdmLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserAdmLog true "用id查询UserAdmLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userAdmLog/findUserAdmLog [get]
export const findUserAdmLog = (params) => {
  return service({
    url: '/adm/findUserAdmLog',
    method: 'get',
    params
  })
}

// @Tags UserAdmLog
// @Summary 分页获取UserAdmLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取UserAdmLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userAdmLog/getUserAdmLogList [get]
export const getUserAdmLogList = (params) => {
  return service({
    url: '/adm/getUserAdmLogList',
    method: 'get',
    params
  })
}
