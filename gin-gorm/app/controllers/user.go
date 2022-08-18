package controllers

import (
	"gin-gorm/app/common/request"
	"gin-gorm/app/common/response"
	"gin-gorm/app/services"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var req request.User
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, -1, err.Error())
		return
	}

	err, _ := services.UserService.AddUser(req)
	if err != nil {
		response.Fail(c, -1, err.Error())
		return
	}

	response.Success(c, "success")
}
