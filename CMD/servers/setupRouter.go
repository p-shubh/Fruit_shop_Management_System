package routers

import (
	"fruit_shop_management_system/CMD/pkg"

	"github.com/gin-gonic/gin"
)

func Setuprouter(c *gin.Engine) {

	// fmt.Println("testing")

	c.POST("/createUser", pkg.CreateUser)
}
