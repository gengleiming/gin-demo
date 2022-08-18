package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username   string `json:"username" binding:"required"`
	Nickname   string `json:"nickname" binding:"required"`
	Email      string `json:"email" binding:"required,email"` // email 校验
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // eqfield 代表两个字段必须相等
	Age        uint8  `json:"age" binding:"gte=1,lte=120"`
	Gender     string `json:"gender" binding:"oneof=female male"`
}

func main() {

	router := gin.Default()

	router.POST("register", Register)

	router.Run(":9999")
}

func Register(c *gin.Context) {
	var r RegisterRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		fmt.Println("register failed")
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	//验证 存储操作省略.....
	fmt.Println("register success")
	c.JSON(http.StatusOK, "successful")
}
