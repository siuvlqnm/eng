package response

import (
	"time"
)

type UserListResponse struct {
	ID           uint   `json:"ID"`
	UserName     string `json:"userName"`
	Phone        string `json:"phone"`
	RegisterPath uint8  `json:"registerPath"`
	UserLevel    uint8  `json:"userLevel"`
	Avatar       string `json:"avatar"`
	// UserUuid     uuid.UUID `json:"userUuid"`
	CreatedAt time.Time `json:"createdAt"`
	JoinTime  time.Time `json:"joinTime"`
}
