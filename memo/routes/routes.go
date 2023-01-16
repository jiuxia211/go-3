package routes

import (
	"paa/memo/api"
	"paa/memo/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	vl := r.Group("api/vl")
	{
		//用户操作
		vl.POST("user/register", api.UserRegister)
		vl.POST("user/login", api.UserLogin)
		authed := vl.Group("/")
		authed.Use(middleware.JwT())
		{
			authed.POST("task", api.CreateTask)
		}
		authed.GET("task/:id", api.ShowTask)
		authed.GET("tasks", api.ListTask)
		authed.PUT("task/:id", api.UpdateTask)
		authed.POST("search", api.SearchTask)
		authed.DELETE("tasks/:id", api.DeleteTask)
	}
	return r
}
