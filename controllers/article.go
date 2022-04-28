package controllers

import (
	"beegodemo/models"
	"github.com/astaxie/beego"
	"time"
)

type ArticleController struct {
	beego.Controller
}

// get: 获取数据， post: 提交数据， put: 修改数据， delete: 删除数据

func (this *ArticleController) Get() {
	this.Data["title"] = "我是标题部分"

	// 解析 html
	this.Data["html"] = "<h3>这是一个<mark>后台渲染</mark>的标题内容</h3>"

	// 模板格式化
	now := time.Now()
	this.Data["now"] = now

	// 获取map对象中的信息
	author := make(map[string]interface{})
	author["name"] = "我是一个作者菌"
	author["created_at"] = 1651048226
	this.Data["author"] = author

	// md5 加密
	this.Data["jiami"] = models.Md5("123456")
	this.TplName = "article.html"
}
