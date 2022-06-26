// 自动生成模板User
package user

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	uuid "github.com/satori/go.uuid"
)

// User 结构体
// 如果含有time.Time 请自行import time包
type User struct {
	global.GVA_MODEL
	UserName     string    `json:"userName" form:"userName" gorm:"column:user_name;comment:用户名;size:191;"`
	Phone        string    `json:"phone" form:"phone" gorm:"column:phone;type:char(11);comment:手机号;"`
	RegisterPath uint8     `json:"registerPath" form:"registerPath" gorm:"column:register_path;default:1;comment:注册途径：1其他，2小程序;"`
	UserLevel    uint8     `json:"userLevel" form:"userLevel" gorm:"column:user_level;default:1;comment:会员等级：1潜在会员，2正式会员，3流失会员;"`
	IsMinor      *bool     `json:"isMinor" form:"isMinor" gorm:"column:is_minor;comment:1成年人、2未成年人;default:1;"`
	Avatar       string    `json:"avatar" form:"avatar" gorm:"column:avatar;comment:头像;size:191;"`
	UserUuid     uuid.UUID `json:"userUuid" form:"userUuid" gorm:"column:user_uuid;comment:用户UUID;size:191;"`
	JoinTime     time.Time `json:"joinTime" form:"joinTime" gorm:"column:join_time;comment:入会时间;"`
}

// TableName User 表名
func (User) TableName() string {
	return "user"
}
