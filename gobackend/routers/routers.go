package routers

import (
	"github.com/gin-gonic/gin"
	"gobackend/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// v1
	v1Group := r.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)

		// 查看所有的代办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
