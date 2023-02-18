package api

import (
	"jiuxia/memo1/service"
	"jiuxia/memo1/utils"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task service.CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&task); err == nil {
		res := task.Create(claim.UserName)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}

}
func UpdateTask(c *gin.Context) {
	var task service.UpdateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&task); err == nil {
		res := task.Update(claim.UserName, c.Param("tid"))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func UpdateAllTask(c *gin.Context) {
	var task service.UpdateAllTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&task); err == nil {
		res := task.UpdateAll(claim.UserName)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func DeleteTask(c *gin.Context) {
	var task service.DeleteTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&task); err == nil {
		res := task.Delete(claim.UserName, c.Param("tid"))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func DeleteAllToDoTask(c *gin.Context) {
	var task service.DeleteTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&task); err == nil {
		res := task.DeleteAllToDo(claim.UserName)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func DeleteAllFinished(c *gin.Context) {
	var task service.DeleteTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&task); err == nil {
		res := task.DeleteAllFinished(claim.UserName)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func ShowTask(c *gin.Context) {
	var task service.ShowTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&task); err == nil {
		res := task.Show(claim.UserName)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func ShowToDoTask(c *gin.Context) {
	var task service.ShowTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&task); err == nil {
		res := task.ShowTodo(claim.UserName)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func ShowFinishedTask(c *gin.Context) {
	var task service.ShowTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&task); err == nil {
		res := task.ShowFinished(claim.UserName)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func SearchTask(c *gin.Context) {
	var task service.SearchTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&task); err == nil {
		res := task.Search(claim.UserName)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
