package servers

import "github.com/gin-gonic/gin"

func Servers() {

	router := gin.Default()

	Setuprouter(router)

	router.Run(":8080")
}
