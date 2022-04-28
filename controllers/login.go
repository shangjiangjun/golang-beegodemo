package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {

	this.TplName = "login.html"
}

func (this *LoginController) DoLogin() {

	// 跳转页面, 重定向： 302 (或者 this.Ctx.Redirect(302, "跳转地址"))
	this.Redirect("/", 302)
}
