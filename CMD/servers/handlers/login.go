package handlers

import (
	"fruit_shop_management_system/CMD/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	reqBody := pkg.Login{}

	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"status": err.Error(),
			"result": reqBody,
		}
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	 

}

