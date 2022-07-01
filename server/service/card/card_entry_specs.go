package card

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	cardReq "github.com/flipped-aurora/gin-vue-admin/server/model/card/request"
)

type CardEntrySpecsService struct {
}

func (s *CardEntrySpecsService) CreateCardEntrySpecs(c cardReq.CardEntryCreate) (err error) {
	err = global.GVA_DB.Create(&c.CardEntrySpecs).Error
	return
}
