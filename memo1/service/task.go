package service

import (
	"jiuxia/memo1/model"
	"jiuxia/memo1/serializer"
	"time"

	"github.com/jinzhu/gorm"
)

type CreateTaskService struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` //0 待办   1已完成
	EndTime int64  `form:"end_at" json:"end_at"`
}
type UpdateTaskService struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status" binding:"required"`
}
type UpdateAllTaskService struct {
	Status int `form:"status" json:"status" bindding:"required"`
}
type DeleteTaskService struct {
}
type ShowTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}
type SearchTaskService struct {
	Info     string `form:"info" json:"info"`
	PageNum  int    `json:"page_num" form:"page_num"`
	PageSize int    `json:"page_size" form:"page_size"`
}

func (service *CreateTaskService) Create(UserName string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("user_name=?", UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "未找到用户",
			}
		}
	}
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		Status:    service.Status,
		StartTime: time.Now().Unix(),
		EndTime:   service.EndTime,
	}
	if err := model.DB.Create(&task).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "创建备忘录失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "创建备忘录成功",
		Data: serializer.Task{
			Tid:       task.ID,
			Title:     task.Title,
			Content:   task.Content,
			Status:    task.Status,
			StartTime: task.StartTime,
			EndTime:   task.EndTime,
		},
	}
}
func (service *UpdateTaskService) Update(UserName string, tid string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("user_name=?", UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "未找到用户",
			}
		}

	}
	var task model.Task
	if err := model.DB.Model(&model.Task{}).First(&task, tid).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "未找到事务",
			}
		}
	}
	if task.Uid == user.ID {
		task.Content = service.Content
		task.Title = service.Title
		task.Status = service.Status
		task.UpdatedAt = time.Now()
		model.DB.Save(&task)
	} else {
		return serializer.Response{
			Status: 400,
			Msg:    "你没有权限修改该事务,因为你不是创建该事务的用户",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "修改成功",
	}

}
func (sercive *UpdateAllTaskService) UpdateAll(UserName string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("user_name=?", UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "未找到用户",
			}
		}

	}
	result := model.DB.Model(&model.Task{}).Where("uid=?", user.ID).Update("status", sercive.Status)
	if result.RowsAffected == 0 {
		return serializer.Response{
			Status: 400,
			Msg:    "未找到事务",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "更新成功",
	}

}
func (service *DeleteTaskService) Delete(UserName, tid string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("user_name=?", UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "未找到用户",
			}
		}

	}
	var task model.Task
	if err := model.DB.Model(&model.Task{}).First(&task, tid).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "未找到事务",
			}
		}
	}
	if task.Uid == user.ID {
		if err := model.DB.Delete(&model.Task{}, tid).Error; err != nil {
			return serializer.Response{
				Status: 400,
				Msg:    "删除失败",
			}
		}
	} else {
		return serializer.Response{
			Status: 400,
			Msg:    "你没有权限删除该事务,因为你不是创建该事务的用户",
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}
func (service *DeleteTaskService) DeleteAllToDo(UserName string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("user_name=?", UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "未找到用户",
			}
		}

	}
	if err := model.DB.Where("uid=? AND status=?", user.ID, 0).Delete(model.Task{}).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "删除失败",
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}
func (service *DeleteTaskService) DeleteAllFinished(UserName string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("user_name=?", UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "未找到用户",
			}
		}

	}
	if err := model.DB.Where("uid=? AND status=?", user.ID, 1).Delete(model.Task{}).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "删除失败",
		}
	}

	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
	}
}
func (service *ShowTaskService) Show(UserName string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("user_name=?", UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "未找到用户",
			}
		}

	}
	var tasks []model.Task
	count := 0
	if service.PageNum == 0 && service.PageSize == 0 {
		service.PageSize = 15
	}
	model.DB.Model(&model.Task{}).Where("uid=? AND status=?", user.ID, 1).Count((&count)).Limit(service.PageSize).
		Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.Response{
		Status: 200,
		Msg:    "数据如上",
		Data: serializer.TaskList{
			Item:  serializer.BuildTasks(tasks),
			Total: uint(count),
		},
	}
}
func (service *ShowTaskService) ShowTodo(UserName string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("user_name=?", UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "未找到用户",
			}
		}

	}
	var tasks []model.Task
	count := 0
	if service.PageNum == 0 && service.PageSize == 0 {
		service.PageSize = 15
	}
	model.DB.Model(&model.Task{}).Where("uid=? AND status=?", user.ID, 1).Count((&count)).Limit(service.PageSize).
		Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.Response{
		Status: 200,
		Msg:    "数据如上",
		Data: serializer.TaskList{
			Item:  serializer.BuildTasks(tasks),
			Total: uint(count),
		},
	}
}
func (service *ShowTaskService) ShowFinished(UserName string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("user_name=?", UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "未找到用户",
			}
		}

	}
	var tasks []model.Task
	count := 0
	if service.PageNum == 0 && service.PageSize == 0 {
		service.PageSize = 15
	}
	model.DB.Model(&model.Task{}).Where("uid=? AND status=?", user.ID, 1).Count((&count)).Limit(service.PageSize).
		Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.Response{
		Status: 200,
		Msg:    "数据如上",
		Data: serializer.TaskList{
			Item:  serializer.BuildTasks(tasks),
			Total: uint(count),
		},
	}
}
func (service *SearchTaskService) Search(UserName string) serializer.Response {
	var user model.User
	if err := model.DB.Model(&model.User{}).Where("user_name=?", UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "未找到用户",
			}
		}

	}
	count := 0
	var tasks []model.Task
	if service.PageNum == 0 && service.PageSize == 0 {
		service.PageSize = 15
	}
	model.DB.Model(&model.Task{}).Where("uid=?", user.ID).Where("title LIKE ? OR content LIKE ?",
		"%"+service.Info+"%", "%"+service.Info+"%").Count((&count)).Limit(service.PageSize).
		Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.Response{
		Status: 200,
		Msg:    "数据如上",
		Data: serializer.TaskList{
			Item:  serializer.BuildTasks(tasks),
			Total: uint(count),
		},
	}
}
