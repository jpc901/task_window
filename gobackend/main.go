package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gobackend/dao"
	"gobackend/models"
	"gobackend/routers"
)

func main() {
	// 创建数据库
	// 链接数据库

	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	// 关闭数据库链接
	defer dao.DB.Close()

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	// 注册路由器
	r := routers.SetupRouter()

	r.Run(":9000")
}
