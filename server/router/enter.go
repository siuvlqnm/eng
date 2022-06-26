package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/card"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/user"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	User    user.RouterGroup
	Card    card.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
