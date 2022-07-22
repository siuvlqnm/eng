package user

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user/response"
)

type UserService struct {
}

// CreateUser 创建User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) CreateUser(user user.User) (err error) {
	err = global.GVA_DB.Create(&user).Error
	return err
}

func (us *UserService) UpdateUserBeforeBuyCard(id uint) (err error) {
	var u *user.User
	upDateMap := make(map[string]interface{}, 2)
	upDateMap["join_time"] = time.Now()
	upDateMap["user_level"] = 2
	err = global.GVA_DB.Model(&u).Where("id = ? and user_level <> 2", id).Updates(upDateMap).Error
	return
}

// DeleteUser 删除User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) DeleteUser(user user.User) (err error) {
	err = global.GVA_DB.Delete(&user).Error
	return err
}

// DeleteUserByIds 批量删除User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) DeleteUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]user.User{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUser 更新User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) UpdateUser(user user.User) (err error) {
	err = global.GVA_DB.Save(&user).Error
	return err
}

// GetUser 根据id获取User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) GetUser(id uint) (user user.User, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&user).Error
	return
}

// GetUserInfoList 分页获取User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) GetUserInfoList(info userReq.UserSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&user.User{})
	var userRsp []response.UserListResponse
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&userRsp).Error
	return userRsp, total, err
}
