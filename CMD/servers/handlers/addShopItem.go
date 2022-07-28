package handlers

import (
	"fmt"
	"fruit_shop_management_system/CMD/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddShopItem(c *gin.Context) {

	reqBody := pkg.Add_Fruits{}

	err := c.Bind(&reqBody)

	if err != nil {

		d, _ := fmt.Println(err)
		res := gin.H{
			"err":    d,
			"status": "request body cant be empty",
			"result": reqBody,
		}

		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return

	}

	if reqBody.Fruits == "" {
		res := gin.H{
			"status": "request body cant be empty",
			"result": reqBody,
		}

		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	result, hence := pkg.InsertFruitsDB(reqBody)

	if !hence {
		resA := gin.H{
			"status": "order failed",
			"result": result,
		}
		c.JSON(http.StatusBadRequest, resA)
	} else {
		resA := gin.H{
			"status": "Order Succesfully created",
			"result": result,
		}
		c.JSON(http.StatusOK, resA)
	}
}
