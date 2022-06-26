package card

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/card"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cardReq "github.com/flipped-aurora/gin-vue-admin/server/model/card/request"
)

type CardEntryService struct {
}

// CreateCardEntry 创建CardEntry记录
// Author [piexlmax](https://github.com/piexlmax)
func (cardEntryService *CardEntryService) CreateCardEntry(cardEntry card.CardEntry) (err error) {
	err = global.GVA_DB.Create(&cardEntry).Error
	return err
}

// DeleteCardEntry 删除CardEntry记录
// Author [piexlmax](https://github.com/piexlmax)
func (cardEntryService *CardEntryService)DeleteCardEntry(cardEntry card.CardEntry) (err error) {
	err = global.GVA_DB.Delete(&cardEntry).Error
	return err
}

// DeleteCardEntryByIds 批量删除CardEntry记录
// Author [piexlmax](https://github.com/piexlmax)
func (cardEntryService *CardEntryService)DeleteCardEntryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]card.CardEntry{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCardEntry 更新CardEntry记录
// Author [piexlmax](https://github.com/piexlmax)
func (cardEntryService *CardEntryService)UpdateCardEntry(cardEntry card.CardEntry) (err error) {
	err = global.GVA_DB.Save(&cardEntry).Error
	return err
}

// GetCardEntry 根据id获取CardEntry记录
// Author [piexlmax](https://github.com/piexlmax)
func (cardEntryService *CardEntryService)GetCardEntry(id uint) (cardEntry card.CardEntry, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cardEntry).Error
	return
}

// GetCardEntryInfoList 分页获取CardEntry记录
// Author [piexlmax](https://github.com/piexlmax)
func (cardEntryService *CardEntryService)GetCardEntryInfoList(info cardReq.CardEntrySearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&card.CardEntry{})
    var cardEntrys []card.CardEntry
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&cardEntrys).Error
	return  cardEntrys, total, err
}
