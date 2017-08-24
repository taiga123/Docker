package controllers

import (
	"github.com/revel/revel"
)

func (c App) SearchResult(name string) revel.Result {
	return c.Render()
}
