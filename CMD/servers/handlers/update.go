package handlers

import (
	"fruit_shop_management_system/CMD/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateItems(c *gin.Context) {

	reqBody := pkg.Add_Fruits{}

	r := c.Bind(&reqBody)

	if r != nil {
		res := gin.H{
			"binding": "error",
			"func":    "create user",
		}
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	if reqBody.Id == 0 || reqBody.Fruits == "" {
		res := gin.H{
			"status": "request body cant be empty",
			"result": reqBody,
		}

		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	result, hence := pkg.UpdateFruits(reqBody)

	if !hence {
		resA := gin.H{
			"status": "order failed",
			"result": result,
		}
		c.JSON(http.StatusBadRequest, resA)
	} else {
		resA := gin.H{
			"status": "Order Succesfully updated",
			"result": result,
		}
		c.JSON(http.StatusOK, resA)
	}

}
