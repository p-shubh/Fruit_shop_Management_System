package servers

import (
	pkg "fruit_shop_management_system/CMD/servers/handlers"

	"github.com/gin-gonic/gin"
)

func Setuprouter(c *gin.Engine) {

	// fmt.Println("testing")

	c.POST("/signupUser/vendor", pkg.CreateUser)
	c.POST("/signupUser", pkg.CreateUser)
	// c.POST("/createShop", pkg.Create_Shop)
	c.POST("/insertfruits", pkg.AddShopItem)
}
