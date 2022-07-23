package user

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/user"
	userReq "github.com/flipped-aurora/gin-vue-admin/server/model/user/request"
)

type UserEntryCardService struct {
}

// CreateUserEntryCard 创建UserEntryCard记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEntryCardService *UserEntryCardService) CreateUserEntryCard(userEntryCard user.UserEntryCard) (err error) {
	err = global.GVA_DB.Create(&userEntryCard).Error
	return err
}

// DeleteUserEntryCard 删除UserEntryCard记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEntryCardService *UserEntryCardService) DeleteUserEntryCard(userEntryCard user.UserEntryCard) (err error) {
	err = global.GVA_DB.Delete(&userEntryCard).Error
	return err
}

// DeleteUserEntryCardByIds 批量删除UserEntryCard记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEntryCardService *UserEntryCardService) DeleteUserEntryCardByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]user.UserEntryCard{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateUserEntryCard 更新UserEntryCard记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEntryCardService *UserEntryCardService) UpdateUserEntryCard(userEntryCard user.UserEntryCard) (err error) {
	err = global.GVA_DB.Save(&userEntryCard).Error
	return err
}

// GetUserEntryCard 根据id获取UserEntryCard记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEntryCardService *UserEntryCardService) GetUserEntryCard(id uint) (userEntryCard user.UserEntryCard, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userEntryCard).Error
	return
}

// GetUserEntryCardInfoList 分页获取UserEntryCard记录
// Author [piexlmax](https://github.com/piexlmax)
func (userEntryCardService *UserEntryCardService) GetUserEntryCardInfoList(info userReq.UserEntryCardSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&user.UserEntryCard{})
	var userEntryCards []user.UserEntryCard
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&userEntryCards).Error
	return userEntryCards, total, err
}

func (s *UserEntryCardService) SuccEntry(usc user.UserEntryCard) {
	if usc.CardType == 2 && usc.IsOpen == 1 {
		return
	}
	if usc.CardType == 1 && usc.IsOpen == 1 && usc.SurplusAmt == 0 {
		return
	}

	if usc.IsOpen == 0 {
		var st = time.Now()
		var et = st
		var d int
		switch usc.DateUnit {
		case 1:
			d = int(usc.ValidPeriod)
		case 2:
			d = int(usc.ValidPeriod) * 30
		case 3:
			d = int(usc.ValidPeriod) * 365
		}
		et = st.AddDate(0, 0, d)
		usc.IsOpen = 1
		usc.StartTime = &st
		usc.EndTime = &et
	} else {
		if usc.SurplusAmt == 1 {
			usc.CardStat = 5
		}
	}
	if usc.CardType == 1 {
		usc.SurplusAmt--
	}
	global.GVA_DB.Where("id = ?", usc.ID).Updates(&usc)
}
