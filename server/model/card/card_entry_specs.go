package card

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type CardEntrySpecs struct {
	global.GVA_MODEL
	CardEntryId uint `gorm:"column:card_entry_id;not null;comment:入场卡ID;"`
	Price       uint `gorm:"column:price;not null;comment:价格;"`
}
