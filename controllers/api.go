package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func (this *ApiController) Get() {
	this.Ctx.WriteString("接收API\n")

	// 获取动态路由
	id := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString("获取接口传值：" + id)
}

// 路由伪静态
func (this *ApiController) CmsRoute() {
	// 获取动态路由
	cmsId := this.Ctx.Input.Param(":id")
	this.Ctx.WriteString("CMS详情数据 \n")
	this.Ctx.WriteString(cmsId)
}
