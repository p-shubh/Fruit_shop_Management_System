package servers

import (
	"fruit_shop_management_system/CMD/servers/handlers"

	"github.com/gin-gonic/gin"
)

func Setuprouter(c *gin.Engine) {

	// fmt.Println("testing")

	c.POST("/signupUser/vendor", handlers.CreateUser)
	c.POST("/signupUser", handlers.CreateUser)
	c.POST("/login", handlers.Login)
	c.GET("/getOrders/:id", handlers.GetOrder)
	c.POST("/insertfruits", handlers.AddShopItem)
	c.PUT("/update", handlers.UpdateItems)
	c.DELETE("/delete", handlers.DeleteItems)
}
