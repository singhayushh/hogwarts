package controllers

import (
	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (i *IndexController) RenderHome(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{
		"title": "Hogwarts",
	})
}
