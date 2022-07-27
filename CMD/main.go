package main

import (
	routers "fruit_shop_management_system/CMD/servers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	routers.Setuprouter(router)

	router.Run(":8080")
}
