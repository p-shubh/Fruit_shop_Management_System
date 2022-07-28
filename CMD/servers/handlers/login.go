package handlers

import (
	"fruit_shop_management_system/CMD/pkg"
	"net/http"
	"strconv"

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

	if pkg.EmailExist(reqBody.Email) {
		reqBody2 := pkg.GetUserByEmail(reqBody.Email)

		c.SetCookie("id", strconv.Itoa(reqBody2.Id), 60*60, "", "", false, false)

		res := gin.H{
			"success": true,
			"message": "sucessfully login",
			"test":    reqBody2,
		}
		c.JSON(http.StatusOK, res)
		return
	}
}
