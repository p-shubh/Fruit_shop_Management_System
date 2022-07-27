package handlers

import (
	"fmt"
	"fruit_shop_management_system/CMD/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOrder(c *gin.Context) {

	ID, err := c.Params.Get("id")
	// fmt.Println("ERRRRRRRRRRRR :",err)

	Id, err2 := strconv.Atoi(ID)

	fmt.Println("err,err2", err, err2)
	if Id == 0 {
		res := gin.H{
			"status": "request body cant be 0 or empty",
			"result": Id,
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	result, hence := pkg.FetchOrders(Id)

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
