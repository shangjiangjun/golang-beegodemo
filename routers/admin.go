package routers

import (
	"beegodemo/admin" // 自定义后台控制器路由
	"github.com/astaxie/beego"
)

// 后台路由
func init() {
	// 后台首页
	beego.Router("/admin", &admin.HomeController{})
}
