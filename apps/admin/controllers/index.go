package controllers

import (
	"seashell/core"
)

type IndexController struct {
	core.BaseController
}

func (c *IndexController) Index() {
	c.Display("admin/index/index")
}

