自定义后台控制器
1. 目录结构

/admin
    +-- 控制器.go

/routers
    +-- router.go


2. 路由文件
```
    package routers

    import (
        "beegodemo/admin" // 自定义后台控制器路由
        "beegodemo/controllers"
        "github.com/astaxie/beego"
    )

    func init() {
        // 后台首页
    	beego.Router("/admin", &admin.HomeController{})
    }

```