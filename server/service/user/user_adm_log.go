package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
)

type UserAdmLogService struct {
}

// CreateUserAdmLog 创建UserAdmLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAdmLogService *UserAdmLogService) CreateUserAdmLog(userAdmLog user.UserAdmLog) (err error) {
	err = global.GVA_DB.Create(&userAdmLog).Error
	return err
}

// DeleteUserAdmLog 删除UserAdmLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAdmLogService *UserAdmLogService) DeleteUserAdmLog(userAdmLog user.UserAdmLog) (err error) {
	err = global.GVA_DB.Delete(&userAdmLog).Error
	return err
}

// DeleteUserAdmLogByIds 批量删除UserAdmLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAdmLogService *UserAdmLogService) DeleteUserAdmLogByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]user.UserAdmLog{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUserAdmLog 更新UserAdmLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAdmLogService *UserAdmLogService) UpdateUserAdmLog(userAdmLog user.UserAdmLog) (err error) {
	err = global.GVA_DB.Save(&userAdmLog).Error
	return err
}

// GetUserAdmLog 根据id获取UserAdmLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAdmLogService *UserAdmLogService) GetUserAdmLog(id uint) (userAdmLog user.UserAdmLog, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userAdmLog).Error
	return
}

// GetUserAdmLogInfoList 分页获取UserAdmLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (userAdmLogService *UserAdmLogService) GetUserAdmLogInfoList(info userReq.UserAdmLogSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&user.UserAdmLog{})
	var userAdmLogs []user.UserAdmLog
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&userAdmLogs).Error
	return userAdmLogs, total, err
}
