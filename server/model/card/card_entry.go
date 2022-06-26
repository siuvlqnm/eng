// 自动生成模板CardEntry
package card

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CardEntry 结构体
// 如果含有time.Time 请自行import time包
type CardEntry struct {
      global.GVA_MODEL
      EntryCardName  string `json:"entryCardName" form:"entryCardName" gorm:"column:entry_card_name;comment:;"`
}


// TableName CardEntry 表名
func (CardEntry) TableName() string {
  return "card_entry"
}

