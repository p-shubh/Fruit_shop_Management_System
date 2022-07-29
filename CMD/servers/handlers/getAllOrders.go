package handlers

import (
	"fruit_shop_management_system/CMD/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get_All_Orders(c *gin.Context) {

	reqBody, bool_value := pkg.GetAllOrders()

	if bool_value {

		res := gin.H{
			"value":  "true",
			"result": reqBody,
		}

		c.JSON(http.StatusOK, res)
	} else if !bool_value {

		res := gin.H{
			"value":  "false",
			"result": reqBody,
		}

		c.JSON(http.StatusBadRequest, res)
	}

}
