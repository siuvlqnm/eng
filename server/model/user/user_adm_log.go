// 自动生成模板UserAdmLog
package user

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// UserAdmLog 结构体
type UserAdmLog struct {
	global.GVA_MODEL
	UserID     uint   `json:"userID" form:"userID" gorm:"column:user_id;not null;"`
	UserName   string `json:"userName" gorm:"column:user_name;not null;comment:用户名;size:191;"`
	Phone      string `json:"phone" gorm:"column:phone;type:char(11);not null;comment:手机号;"`
	UserCardID uint   `json:"userCardID" form:"userCardID" gorm:"column:user_card_id;not null;"`
	CardName   string `json:"cardName" gorm:"column:card_name;not null;comment:入场卡名称;"`
	EntryType  uint8  `json:"entryType" form:"entryType" gorm:"column:entry_type;not null;default:1;comment:入场类型：1会员入场，2预约到访;"`
	// EntryTime time.Time  `json:"entryTime" gorm:"column:entry_time;type:timestamp;comment:入场时间;"`
	LeaveTime *time.Time `json:"leaveTime" gorm:"column:leave_time;type:timestamp;comment:离场时间;"`
	TotalTime uint16     `json:"totalTime" gorm:"column:total_time;not null;default:0;comment:在场时间;"`
	DeductNum uint8      `json:"deductNum" form:"deductNum" gorm:"column:deduct_num;not null;default:0;comment:扣除次数;"`
}

// TableName UserAdmLog 表名
func (UserAdmLog) TableName() string {
	return "user_admission_log"
}
