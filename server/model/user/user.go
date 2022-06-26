// 自动生成模板User
package user

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// User 结构体
// 如果含有time.Time 请自行import time包
type User struct {
      global.GVA_MODEL
      UserName  string `json:"userName" form:"userName" gorm:"column:user_name;comment:;"`
}


// TableName User 表名
func (User) TableName() string {
  return "user"
}

