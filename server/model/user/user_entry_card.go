// 自动生成模板UserEntryCard
package user

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	uuid "github.com/satori/go.uuid"
)

// UserEntryCard 结构体
type UserEntryCard struct {
	global.GVA_MODEL
	UserID         uint      `json:"userID" form:"userID" gorm:"column:user_id;not null;"`
	UserName       string    `json:"userName" form:"userName" gorm:"column:user_name;not null;comment:用户名;size:191;"`
	Phone          string    `json:"phone" form:"phone" gorm:"column:phone;type:char(11);not null;comment:手机号;"`
	CardID         uint      `json:"cardID" form:"cardID" gorm:"column:card_id;not null;"`
	AuditStat      uint8     `json:"aduitStat" gorm:"column:aduit_status;not null;default:0;comment:审核状态：0未审核，1审核通过，2审核未通过;"`
	CardName       string    `json:"cardName" gorm:"column:card_name;not null;comment:入场卡名称;"`
	CardType       uint8     `json:"cardType" gorm:"column:card_type;not null;comment:入场卡类型：1次卡，2期限卡;"`
	CardStat       uint8     `json:"cardStat" gorm:"column:card_status;not null;default:1;comment:卡状态：1有效未冻结，2有效冻结，3失效退款，4失效过期，5失效耗尽，6失效转让，7失效升级;"`
	IsOpen         uint8     `json:"isOpen" gorm:"column:is_open;not null;default:0;comment:是否开卡：0未开卡，1已开卡;"`
	IsRefund       uint8     `json:"isRefund" gorm:"column:is_refund;not null;default:0;comment:是否退款：0未退款，1已退款，2，3;"`
	InitAmt        uint16    `json:"initAmt" gorm:"column:init_amount;not null;default:0;comment:初始化数量;"`
	TotalAmt       uint16    `json:"totalAmt" gorm:"column:total_amount;not null;default:0;comment:总数量;"`
	GiftAmt        uint16    `json:"giftAmt" form:"giftAmt" gorm:"column:gift_amount;not null;default:0;comment:赠送数量;"`
	SurplusAmt     uint16    `json:"surplusAmt" gorm:"column:surplus_amount;not null;default:0;comment:剩余数量;"`
	ContractNumber uuid.UUID `json:"contractNumber" gorm:"column:contract_number;not null;size:191;comment:合同编号;"`
	TotalPrice     uint      `json:"totalPrice" gorm:"column:total_price;not null;default:0;comment:原价格;"`
	PayPrice       uint      `json:"payPrice" form:"payPrice" gorm:"column:pay_price;not null;default:0;comment:实际支付价格;"`
	ValidTime      uint16    `json:"validTime" gorm:"column:valid_time;not null;default:0;comment:有效次数;"`
	ValidPeriod    uint16    `json:"validPeriod" gorm:"column:valid_period;not null;default:1;comment:有效期;"`
	DateUnit       uint8     `json:"dateUnit" gorm:"column:date_unit;not null;default:1;comment:日期单位：1天，2月，3年;"`
	StartTime      time.Time `json:"startTime" gorm:"column:start_time;type:timestamp;not null;comment:激活时间;"`
	EndTime        time.Time `json:"endTime" gorm:"column:end_time;type:timestamp;not null;comment:到期时间;"`
	Remark         string    `json:"remark" form:"remark" gorm:"column:remark;not null;comment:备注;"`
}

// TableName UserEntryCard 表名
func (UserEntryCard) TableName() string {
	return "user_entry_card"
}
