package card

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type CardEntrySpecs struct {
	global.GVA_MODEL
	CardEntryId uint   `gorm:"column:card_entry_id;not null;comment:入场卡ID;"`
	Price       uint16 `json:"price" form:"price" gorm:"column:price;not null;default:0;comment:价格;"`
	ValidTime   uint16 `json:"validTime" form:"validTime" gorm:"column:valid_time;not null;default:0;comment:有效次数;"`
	ValidPeriod uint16 `json:"validPeriod" form:"validPeriod" gorm:"column:valid_period;not null;default:1;comment:有效期;"`
	DateUnit    uint8  `json:"dateUnit" form:"dateUnit" gorm:"column:date_unit;not null;default:1;comment:日期单位：1天，2月，3年;"`
	FreezeAllow uint16 `json:"freezeAllow" form:"freezeAllow" gorm:"column:freeze_allow;not null;default:0;comment:允许冻结天数;"`
	GiftAllow   uint16 `json:"giftAllow" form:"giftAllow" gorm:"column:gift_allow;not null;default:0;comment:允许赠送天数;"`
}

func (CardEntrySpecs) TableName() string {
	return "card_entry_specs"
}
