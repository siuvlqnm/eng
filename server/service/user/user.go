package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
)

type UserService struct {
}

// CreateUser 创建User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService) CreateUser(user user.User) (err error) {
	err = global.GVA_DB.Create(&user).Error
	return err
}

// DeleteUser 删除User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService)DeleteUser(user user.User) (err error) {
	err = global.GVA_DB.Delete(&user).Error
	return err
}

// DeleteUserByIds 批量删除User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService)DeleteUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]user.User{},"id in ?",ids.Ids).Error
	return err
}

// UpdateUser 更新User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService)UpdateUser(user user.User) (err error) {
	err = global.GVA_DB.Save(&user).Error
	return err
}

// GetUser 根据id获取User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService)GetUser(id uint) (user user.User, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&user).Error
	return
}

// GetUserInfoList 分页获取User记录
// Author [piexlmax](https://github.com/piexlmax)
func (userService *UserService)GetUserInfoList(info userReq.UserSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&user.User{})
    var users []user.User
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&users).Error
	return  users, total, err
}
