package handlers

import (
	"fruit_shop_management_system/CMD/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteItems(c *gin.Context) {
	reqBody := pkg.Id{}

	c.Bind(&reqBody)

	if reqBody.Id == 0 {
		res := gin.H{
			"status": "request body cant be empty",
			"result": reqBody,
		}

		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	result, hence := pkg.DeleteFruitItems(reqBody.Id)

	if !hence {
		resA := gin.H{
			"status": "order failed",
			"result": result,
		}
		c.JSON(http.StatusBadRequest, resA)
	} else {
		resA := gin.H{
			"status": "Order Succesfully deleted",
			"result": result,
		}
		c.JSON(http.StatusOK, resA)
	}
}
