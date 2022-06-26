package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/card"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CardEntrySearch struct{
    card.CardEntry
    request.PageInfo
}
