package admin

import (
	"beegodemo/models/database"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"regexp"
	"strconv"
)

type UsersController struct {
	beego.Controller
}

func (this *UsersController) Get() {
	this.Data["title"] = "用户管理列表"

	var fields []string
	var sortby = []string{
		"id",
	}
	var order = []string{
		"desc",
	}
	var query = make(map[string]string)
	user_lists, err := database.GetAllUsers(query, fields, sortby, order, 0, 0)
	if err == nil {
		this.Data["users"] = user_lists
	} else {
		this.Data["users"] = make(map[string]string)
	}

	this.TplName = "admin/users/index.html"
}

// 到新增页面
func (this *UsersController) NewUser() {
	this.TplName = "admin/users/form.html"
}

func (this *UsersController) AddUser() {
	user := database.Users{
		Name:   this.GetString("name"),
		Mobile: this.GetString("mobile"),
	}

	response := make(map[string]interface{})
	// 执行编辑功能
	if _, err := database.AddUser(&user); err == nil {
		response["status"] = true
		response["code"] = 200
		response["message"] = "新建成功"
		this.Data["json"] = response
		this.ServeJSON()
		this.StopRun()
	}
	response["status"] = false
	response["code"] = 500
	response["message"] = "新建失败"
	this.Data["json"] = response
	this.ServeJSON()
	this.StopRun()
}

func (this *UsersController) GetInfo() {
	// 动态参数
	//id, _ := this.GetInt64(":id")
	id := this.Ctx.Input.Param(":id")
	//this.Ctx.WriteString("获取接口传值：" + id)
	id64, _ := strconv.ParseInt(id, 10, 64)
	//fmt.Printf("%v ---- %T \n", id64, id64)
	info, err := database.GetUserById(id64)
	if err != nil {
		this.Data["message"] = "参数错误"
		this.TplName = "error.html"
		this.StopRun()
	}
	this.Data["info"] = info
	this.TplName = "admin/users/form.html"
}

func (this *UsersController) EditUser() {
	Id, _ := this.GetInt64("id")
	user := database.Users{
		Id:     Id,
		Name:   this.GetString("name"),
		Mobile: this.GetString("mobile"),
	}

	if err := this.ParseForm(&user); err != nil {
		this.Ctx.WriteString("数据提交失败")
		return
	} else {
		// 打印结构体内的所有信息
		//fmt.Printf("%#v \n", user)

		valid := validation.Validation{}
		// params := this.ParseForm()
		// Required 不为空，即各个类型要求不为其零值
		res := valid.Required(user.Name, "name")
		fmt.Printf("%#v \n", res)
		if !res.Ok {
			//this.Ctx.WriteString(fmt.Sprintln("Name用户名 : ", res.Error.Key, res.Error.Message))
			this.Data["message"] = res.Error.Key + ", " + res.Error.Message
			this.TplName = "error.html"
			return
		}

		res = valid.Mobile(user.Mobile, "mobile")
		if !res.Ok {
			//this.Ctx.WriteString(fmt.Sprintln("Mobile手机号 : ", res.Error.Key, res.Error.Message))
			this.Data["message"] = res.Error.Key + ", " + res.Error.Message
			this.TplName = "error.html"
			return
		}
		// 执行编辑功能
		if result := database.UpdateUserById(&user); result == nil {
			this.Data["message"] = "操作成功！"
			this.Data["linkUrl"] = "/admin/users"
			this.Data["timeout"] = 3
			this.TplName = "success.html"
			return
		}
		this.Data["message"] = "编辑失败"
		this.TplName = "error.html"
	}
}

// 表单验证模板
func (this *UsersController) Index() {
	valid := validation.Validation{}
	// Required 不为空，即各个类型要求不为其零值
	res := valid.Required(nil, "name")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Required 不为空 : ", res.Error.Key, res.Error.Message))
	}

	// Min(min int) 最小值，有效类型：int，其他类型都将不能通过验证
	res = valid.Min(16, 18, "min_age")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Min(min int) 最小值 : ", res.Error.Key, res.Error.Message))
	}
	// Max(max int) 最大值，有效类型：int，其他类型都将不能通过验证
	res = valid.Max(20, 19, "max_age")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Max(max int) 最大值 : ", res.Error.Key, res.Error.Message))
	}
	// Range(min, max int) 数值的范围，有效类型：int，他类型都将不能通过验证
	res = valid.Range(nil, 16, 18, "range_age")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Range(min, max int) 数值的范围 : ", res.Error.Key, res.Error.Message))
	}
	// MinSize(min int) 最小长度，有效类型：string slice，其他类型都将不能通过验证
	res = valid.MinSize(123, 5, "min_size")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("MinSize(min int) 最小长度 : ", res.Error.Key, res.Error.Message))
	}
	// MaxSize(max int) 最大长度，有效类型：string slice，其他类型都将不能通过验证
	res = valid.MaxSize(123, 2, "max_size")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("MaxSize(max int) 最大长度 : ", res.Error.Key, res.Error.Message))
	}
	// Length(length int) 指定长度，有效类型：string slice，其他类型都将不能通过验证
	res = valid.Length(0, 1, "length")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Length(length int) 指定长度 : ", res.Error.Key, res.Error.Message))
	}
	// Alpha alpha字符，有效类型：string，其他类型都将不能通过验证
	// res = valid.Alpha("", "alpha")
	res = valid.Alpha(nil, "alpha")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Alpha alpha字符 : ", res.Error.Key, res.Error.Message))
	}
	// Numeric 数字，有效类型：string，其他类型都将不能通过验证
	// res = valid.Numeric("2", "numeric")
	res = valid.Numeric(2, "numeric")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Numeric 数字 : ", res.Error.Key, res.Error.Message))
	}
	// AlphaNumeric alpha 字符或数字，有效类型：string，其他类型都将不能通过验证
	res = valid.AlphaNumeric(nil, "AlphaNumeric")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("AlphaNumeric alpha 字符或数字 : ", res.Error.Key, res.Error.Message))
	}
	// Match(pattern string) 正则匹配，有效类型：string，其他类型都将被转成字符串再匹配(fmt.Sprintf(“%v”, obj).Match)
	// res = valid.Match("123456789", regexp.MustCompile(`^(\-|\+)?\d+(\.\d+)?$`), "Match")
	res = valid.Match("abc", regexp.MustCompile(`^(\-|\+)?\d+(\.\d+)?$`), "Match")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Match(pattern string) 正则匹配 : ", res.Error.Key, res.Error.Message))
	}
	// AlphaDash alpha字符或数字或横杠-_，有效类型：string，其他类型都将不能通过验证
	res = valid.AlphaDash(nil, "AlphaDash")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("AlphaDash alpha字符或数字或横杠-_，有效类型 : ", res.Error.Key, res.Error.Message))
	}
	// Email邮箱格式，有效类型：string，其他类型都将不能通过验证
	// res = valid.Email("123456@qq.com", "email")
	res = valid.Email("123456qq.com", "email")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Email邮箱格式 : ", res.Error.Key, res.Error.Message))
	}
	// IP IP格式，目前只支持IPv4格式验证，有效类型：string，其他类型都将不能通过验证
	// res = valid.IP("192.168.0.1", "ip")
	res = valid.IP("192.168.300.1", "ip")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("IP格式 : ", res.Error.Key, res.Error.Message))
	}
	// Base64 base64编码，有效类型：string，其他类型都将不能通过验证
	// res = valid.Base64(base64.StdEncoding.EncodeToString([]byte("abc")), "base64")
	res = valid.Base64(nil, "base64")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("base64编码 : ", res.Error.Key, res.Error.Message))
	}
	// Mobile手机号，有效类型：string，其他类型都将不能通过验证
	// res = valid.Mobile("+8615621628869", "mobile")
	// res = valid.Mobile("15621628869", "mobile")
	// res = valid.Mobile(15621628869, "mobile")
	res = valid.Mobile("+861528869", "mobile")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Mobile手机号 : ", res.Error.Key, res.Error.Message))
	}
	// Tel固定电话号，有效类型：string，其他类型都将不能通过验证
	// res = valid.Tel("010-7700008", "tel")
	res = valid.Tel("15621628869", "tel")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Tel固定电话号 : ", res.Error.Key, res.Error.Message))
	}
	// Phone手机号或固定电话号，有效类型：string，其他类型都将不能通过验证
	res = valid.Phone("110", "phone")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("Phone手机号或固定电话号 : ", res.Error.Key, res.Error.Message))
	}
	// ZipCode邮政编码，有效类型：string，其他类型都将不能通过验证
	// res = valid.ZipCode("100000", "zipcode")
	res = valid.ZipCode("000000", "zipcode")
	if !res.Ok {
		this.Ctx.WriteString(fmt.Sprintln("ZipCode邮政编码 : ", res.Error.Key, res.Error.Message))
	}
}

// 删除数据
func (this *UsersController) DeleteUser() {
	response := make(map[string]interface{})

	id, _ := this.GetInt64("id")
	if err := database.DeleteUser(id); err == nil {
		response["status"] = true
		response["code"] = 200
		response["message"] = "删除成功"
		this.Data["json"] = response
		this.ServeJSON()
		this.StopRun()
	}
	response["status"] = false
	response["code"] = 500
	response["message"] = "删除失败"
	this.Data["json"] = response
	this.ServeJSON()
	this.StopRun()

}
