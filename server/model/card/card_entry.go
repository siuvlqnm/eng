// 自动生成模板CardEntry
package card

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CardEntry 结构体
// 如果含有time.Time 请自行import time包
type CardEntry struct {
	global.GVA_MODEL
	CardName      string    `json:"cardName" form:"cardName" gorm:"column:card_name;not null;comment:入场卡名称;"`
	CardType      uint8     `json:"cardType" form:"cardType" gorm:"column:card_type;not null;default:1;comment:入场卡类型：1次卡，2期限卡;"`
	SuptUserNr    uint8     `json:"suptUserNr" form:"suptUserNr" gorm:"column:support_user_number;not null;default:1;comment:支持入场人数;"`
	PriceGradient string    `json:"priceGradient" form:"priceGradient" gorm:"column:price_gradient;not null;comment:价格区间;"`
	TimeGradient  string    `json:"timeGradient" form:"timeGradient" gorm:"column:time_gradient;not null;comment:时间区间;"`
	CardStat      uint8     `json:"cardStat" form:"cardStat" gorm:"column:card_status;not null;default:1;comment:上架状态：1已上架可售，2已下架可售，3已下架停售;"`
	SaleType      uint8     `json:"saleType" form:"saleType" gorm:"column:sale_type;not null;default:1;comment:售卖类型：1永久售卖，2限时售卖;"`
	SaleStartTime time.Time `json:"saleStartTime" form:"saleStartTime" gorm:"column:sale_start_time;type:timestamp;comment:售卖开始时间;"`
	SaleEndTime   time.Time `json:"saleEndTime" form:"saleEndTime" gorm:"column:sale_end_time;type:timestamp;comment:售卖结束时间;"`
	IsTransfer    uint8     `json:"isTransfer" form:"isTransfer" gorm:"column:is_transfer;not null;default:0;comment:是否支持转让;"`
	TransferFee   uint      `json:"transferFee" form:"transferFee" gorm:"column:transfer_fee;not null;comment:转让费用;"`
	TransferUnit  uint8     `json:"transferUnit" form:"transferUnit" gorm:"column:transfer_unit;not null;default:1;comment:转让计算单位：1元，2百分比;"`
	SpecialNote   string    `json:"specialNote" form:"specialNote" gorm:"column:special_note;not null;comment:特别说明;"`
	CardIntro     string    `json:"cardIntro" form:"cardIntro" gorm:"column:card_introduction;not null;comment:卡介绍"`
	IntIntro      string    `json:"intIntro" form:"intIntro" gorm:"column:internal_introduction;not null;comment:内部说明"`
}

// TableName CardEntry 表名
func (CardEntry) TableName() string {
	return "card_entry"
}
