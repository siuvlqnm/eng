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
	CardName   string `json:"cardName" form:"cardName" gorm:"column:card_name;not null;comment:入场卡名称;"`
	CardType   uint8  `json:"cardType" form:"cardType" gorm:"column:card_type;not null;default:1;comment:入场卡类型：1次卡，2期限卡;"`
	SuptUserNr uint8  `json:"suptUserNr" form:"suptUserNr" gorm:"column:support_user_number;not null;default:1;comment:支持入场人数;"`

	CardStat      uint8     `json:"cardStat" form:"cardStat" gorm:"column:card_status;not null;default:1;comment:上架状态：1已上架可售，2已下架可售，3已下架停售;"`
	SaleType      uint8     `json:"saleType" form:"saleType" gorm:"column:sale_type;not null;default:1;comment:售卖类型：1永久售卖，2限时售卖;"`
	SaleStartTime time.Time `json:"saleStartTime" form:"saleStartTime" gorm:"column:sale_time_time;type:timestamp;not null;comment:售卖开始时间;"`
	SaleEndTime   time.Time `json:"saleEndTime" form:"saleEndTime" gorm:"column:sale_end_time;type:timestamp;not null;comment:售卖结束时间;"`
}

// TableName CardEntry 表名
func (CardEntry) TableName() string {
	return "card_entry"
}
