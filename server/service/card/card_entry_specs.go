package card

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/card"
	cardReq "github.com/flipped-aurora/gin-vue-admin/server/model/card/request"
)

type CardEntrySpecsService struct {
}

func (s *CardEntrySpecsService) CreateCardEntrySpecs(c cardReq.CardEntryCreate) (err error) {
	err = global.GVA_DB.Create(&c.CardEntrySpecs).Error
	return
}

func (s *CardEntrySpecsService) GetCardEntrySpecs(id uint) (ces card.CardEntrySpecs, err error) {
	err = global.GVA_DB.Where("card_entry_id = ?", id).First(&ces).Error
	return
}
