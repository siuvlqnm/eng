package card

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type CardEntrySpecs struct {
	global.GVA_MODEL
	CardEntryId uint   `gorm:"column:card_entry_id;not null;comment:入场卡ID;"`
	Price       uint   `gorm:"column:price;not null;default:0;comment:价格;"`
	ValidTime   uint   `gorm:"column:valid_time;not null;default:0;comment:有效次数;"`
	ValidPeriod uint   `gorm:"column:valid_period;not null;default;1;comment:有效周期;"`
	DateType    string `gorm:"column:date_type;not null;default:1;comment:日期类型：1天，2月，3年;"`
	FreezeAllow uint   `gorm:"column:freeze_allow;not null;default:0;comment:允许冻结天数;"`
	GiftAllow   uint   `gorm:"column:gift_allow;not null;default:0;comment:允许赠送天数;"`
}
