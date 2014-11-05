package controllers

import (
	"github.com/revel/revel"
)

func init() {

	revel.InterceptMethod((*AppController).checkAuth, revel.BEFORE)
}
