package api

import (
	"jiuxia/memo1/service"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var useRegister service.UseService
	if err := c.ShouldBind(&useRegister); err == nil {
		res := useRegister.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func UserLogin(c *gin.Context) {
	var userLogin service.UseService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
