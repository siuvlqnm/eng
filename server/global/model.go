package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gorm:"primarykey;not null;"`           // 主键ID
	CreatedAt time.Time      `gorm:"type:timestamp;not null;"`       // 创建时间
	UpdatedAt time.Time      `gorm:"type:timestamp;not null;"`       // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index;type:timestamp;" json:"-"` // 删除时间
}
