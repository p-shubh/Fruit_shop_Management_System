package handlers

import (
	"fruit_shop_management_system/CMD/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get_All_Orders(c *gin.Context) {

	result, bool_value := pkg.GetAllOrdersService()

	if bool_value {

		res := gin.H{
			"value":  "true",
			"result": result,
		}

		c.JSON(http.StatusOK, res)
	} else if !bool_value {

		res := gin.H{
			"value":  "false",
			"result": result,
		}

		c.JSON(http.StatusBadRequest, res)
	}

}
