package main

import (
	"beegodemo/models"
	_ "beegodemo/routers"
	"github.com/astaxie/beego"
)

func main() {
	// 引入自定义函数 models 文件夹中
	beego.AddFuncMap("unixToDate", models.UnixToDate)
	beego.Run()
}
