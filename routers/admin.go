package routers

import (
	"beegodemo/admin" // 自定义后台控制器路由
	"github.com/astaxie/beego"
)

// 后台路由
func init() {
	// 后台首页
	beego.Router("/admin", &admin.HomeController{})
	beego.Router("/admin/users", &admin.UsersController{})
	beego.Router("/admin/users/new", &admin.UsersController{}, "get:NewUser")
	beego.Router("/admin/users/add", &admin.UsersController{}, "post:AddUser")
	beego.Router("/admin/users/:id", &admin.UsersController{}, "get:GetInfo")
	beego.Router("/admin/users/edit", &admin.UsersController{}, "put:EditUser")
	beego.Router("/admin/users/delete", &admin.UsersController{}, "delete:DeleteUser")
}
