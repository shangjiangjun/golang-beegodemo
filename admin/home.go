package admin

import (
	"beegodemo/models/database"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

// 获取数据库 用户信息
func (this *HomeController) Get() {
	response := make(map[string]interface{})

	// 1. 获取链接对象
	//o := orm.NewOrm()

	// 2. 查询所有数据
	/*var user_lists []database.Users

	_, err := o.QueryTable("users").All(&user_lists)
	*/
	var fields []string
	var sortby = []string{
		"id",
	}
	var order = []string{
		"desc",
	}
	var query = make(map[string]string)
	user_lists, err := database.GetAllUsers(query, fields, sortby, order, 0, 0)
	if err != nil {
		response["status"] = false
		response["code"] = 500
		response["message"] = err.Error()
		this.Data["json"] = response
		this.ServeJSON()
		this.StopRun()
	}
	/*// 1. 获取链接对象
	o := orm.NewOrm()
	var articles []database.Articles
	// 条件查询： Filter()
	_, err2 := o.QueryTable("articles").Filter("CategoryId", 4).All(&articles)
	if err2 != nil {
		response["status"] = false
		response["code"] = 500
		response["message"] = err.Error()
		this.Data["json"] = response
		this.ServeJSON()
		this.StopRun()
	}*/

	response["status"] = true
	response["code"] = 200
	response["message"] = ""
	data := make(map[string]interface{})
	data["users"] = user_lists
	//data["articles"] = articles
	response["data"] = data
	this.Data["json"] = response
	this.ServeJSON()
}
