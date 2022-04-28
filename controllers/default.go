package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type MainController struct {
	beego.Controller
}

type Article struct {
	Title   string
	Content string
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	// 模板中绑定结构体数据
	article := Article{
		Title:   "结构体标题",
		Content: "ajhgadf",
	}
	c.Data["article"] = article

	// 模板遍历切片
	c.Data["sliceList"] = []string{"php", "go", "java"}

	// 循环遍历map
	userinfo := make(map[string]interface{})
	userinfo["username"] = "用户名称"
	userinfo["age"] = 23
	c.Data["userinfo"] = userinfo

	// 结构体 类型的切片
	c.Data["articleList"] = []Article{
		{
			Title:   "新闻1",
			Content: "新闻内容1",
		},
		{
			Title:   "新闻2",
			Content: "新闻内容2",
		},
		{
			Title:   "新闻3",
			Content: "新闻内容3",
		},
		{
			Title:   "新闻4",
			Content: "新闻内容4",
		},
	}

	// 匿名结构体
	/*
		[]string{"1", "2", "3"}

		[]struct {
			Title string
		}{
			{Title: "匿名结构体内容1"},{Title: "匿名结构体内容1"},{Title: "匿名结构体内容1"},
		}
	*/
	c.Data["cmsList"] = []struct {
		Title string
	}{
		{Title: "匿名结构体内容1"}, {Title: "匿名结构体内容2"}, {Title: "匿名结构体内容3"},
	}
	c.TplName = "index.html"
}

// Get请求，必须首字母大写
func (this *MainController) ParamGet() {
	//id := this.GetString("id")
	// 判断类型
	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.WriteString("传入的参数id值必须是int类型")
		return
	}
	// 如果还用 ”+“ 会报错，因为id为数值型, strconv.Itoa() 数值转字符串
	this.Ctx.WriteString("接收到的数据是：" + strconv.Itoa(id))

	// 打印在终端命令栏中
	beego.Info(id)
	fmt.Printf("%v --- %T \n", id, id)
}

// 定义结构体 form: 接收的数据传参名， json:转json的key
type User struct {
	Username string   `form:"username" json:"username"`
	Password string   `form:"password" json:"password"`
	Hobby    []string `form:"hobby" json:"hobby"`
}

func (this *MainController) ParamPost() {
	user := User{}

	if err := this.ParseForm(&user); err != nil {
		this.Ctx.WriteString("数据提交失败")
		return
	}

	// 打印结构体内的所有信息
	fmt.Printf("%#v \n", user)
	this.Ctx.WriteString("解析数据完成\n")

	this.Data["json"] = user
	this.ServeJSON()
}

/*
func (this *MainController) ParamPost() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.WriteString("传入的参数id值必须是int类型")
		return
	}
	username := this.GetString("username")
	password := this.GetString("password")
	hobby := this.GetStrings("hobby")
	fmt.Printf("值： %v ------ 类型： %T \n", hobby, hobby)
	this.Ctx.WriteString("用户提交数据为：" + strconv.Itoa(id) + username + "," + password)
}
*/

// 返回json数据
func (this *MainController) BackJson() {
	u := User{
		Username: "站啊啊水水",
		Password: "123123",
		Hobby:    []string{"1", "2"},
	}

	this.Data["json"] = u
	this.ServeJSON()
}
