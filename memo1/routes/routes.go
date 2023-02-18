package routes

import (
	"jiuxia/memo1/api"
	"jiuxia/memo1/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authorized := v1.Group("/")
		authorized.Use(middleware.JwT())
		{
			authorized.POST("task/create", api.CreateTask)
			authorized.PUT("task/:tid", api.UpdateTask)
			authorized.PUT("task/update", api.UpdateAllTask)
			authorized.DELETE("task/:tid", api.DeleteTask)
			authorized.DELETE("task/delete/todo", api.DeleteAllToDoTask)
			authorized.DELETE("task/delete/finished", api.DeleteAllFinished)
			authorized.GET("task/show", api.ShowTask)
			authorized.GET("task/show/todo", api.ShowToDoTask)
			authorized.GET("task/show/finished", api.ShowFinishedTask)
			authorized.POST("search", api.SearchTask)
		}

	}
	return r
}
