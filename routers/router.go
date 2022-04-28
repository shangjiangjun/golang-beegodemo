package routers

import (
	"beegodemo/controllers"
	"github.com/astaxie/beego"
)

// 前端路由
func init() {
	beego.Router("/", &controllers.MainController{})

	// get 请求
	beego.Router("/param/get", &controllers.MainController{}, "get:ParamGet")

	// 表单提交： /param/post
	beego.Router("/param/post", &controllers.MainController{}, "post:ParamPost")

	// json 输出示例
	beego.Router("/param/json", &controllers.MainController{}, "get:BackJson")

	// get: 获取数据， post: 提交数据， put: 修改数据， delete: 删除数据
	// 如果路由报错，则要查看控制器中的名称及控制器名的对应上是否正确
	beego.Router("/methods", &controllers.MethodController{})
	beego.Router("/methods/add", &controllers.MethodController{}, "post:DealAdd")
	beego.Router("/methods/edit", &controllers.MethodController{}, "put:DoEdit")
	beego.Router("/methods/delete", &controllers.MethodController{}, "delete:DoDelete")

	// xml数据处理
	beego.Router("/methods/xml", &controllers.MethodController{}, "post:Xml")

	// 动态路由， 请求地址： /api/123
	beego.Router("/api/:id", &controllers.ApiController{})

	// 正则路由: 伪静态使用 请求地址： /cms_123.html
	beego.Router("/cms_:id([0-9]+).html", &controllers.ApiController{}, "get:CmsRoute")

	// 登录页面
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/do-login", &controllers.LoginController{}, "post:DoLogin")

	// 公共页面调用
	beego.Router("/article", &controllers.ArticleController{})

}
