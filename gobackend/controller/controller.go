package controller

import (
	"github.com/gin-gonic/gin"
	"gobackend/models"
	"net/http"
)

/*
 url --> controller --> logic   --> model
请求 -->  控制器     --> 业务逻辑  --> 模型的增删改查
*/

func CreateTodo(c *gin.Context) {
	// 前端页面填写代办事项，点击提交，会发送请求到这里
	// 1.从请求中吧数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	// 2.存入数据库
	// 3.返回响应
	err := models.CreateATodo(&todo)
	if err != nil {
		// 存入数据库失败
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
	} else {
		// 存入数据成功
		c.JSON(http.StatusOK, todo)
	}
}
func GetTodoList(c *gin.Context) {
	// 查看表里的所有数据
	todoList, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}

}

func UpdateATodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := models.GetATodo(id)
	c.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id := c.Param("id")
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
