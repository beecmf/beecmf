package controllers

import (
	"github.com/beecmf/beecmf"
)

type IndexController struct {
	beecmf.BaseController
}

func (c *IndexController) Index() {
	c.Display("admin/index/index")
}

