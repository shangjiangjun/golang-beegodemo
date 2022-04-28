package controllers

import (
	"encoding/xml"
	"github.com/astaxie/beego"
)

type MethodController struct {
	beego.Controller
}

// get: 获取数据， post: 提交数据， put: 修改数据， delete: 删除数据

func (this *MethodController) Get() {
	this.Ctx.WriteString("执行查询动作")
}

func (this *MethodController) DealAdd() {
	this.Ctx.WriteString("执行增加动作")
}

func (this *MethodController) DoEdit() {
	title := this.GetString("string")
	this.Ctx.WriteString("执行修改动作" + title)
}

func (this *MethodController) DoDelete() {
	this.Ctx.WriteString("执行删除动作")
}

// 数据xml
type Product struct {
	Title   string `from:"title" xml:"title"` // 注意xml 的映射
	Content string `from:"content" xml:"content"`
}

/*
<?xml version="1.0" encoding="UTF-8">
<article>
	<title type="string">我是标题</title>
	<content type="string">我是内容部分</content>
</article>
*/
// 接收Post 传递过来的xml数据
// 需要在配置文件中配置 copyrequestbody=true, 需要重启
func (this *MethodController) Xml() {
	//str := string(this.Ctx.Input.RequestBody)
	//beego.Info(str)
	//this.Ctx.WriteString(str)

	p := Product{}
	// 解析
	var err error
	if e := xml.Unmarshal(this.Ctx.Input.RequestBody, &p); e != nil {
		this.Data["json"] = err // err.Error() 无指针
		this.ServeJSON()
	} else {
		this.Data["json"] = p
		this.ServeJSON()
	}

}
