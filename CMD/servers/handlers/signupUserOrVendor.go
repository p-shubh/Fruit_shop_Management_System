package handlers

import (
	"fruit_shop_management_system/CMD/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

var user_type int

func CreateUser(c *gin.Context) {

	reqBody := pkg.Create_Shop_Account{}

	c.Bind(&reqBody)

	if c.Request.URL.Path == "/signupUser/vendor" {
		user_type = 1
	} else {
		user_type = 2
	}

	// if r == nil {
	// 	res := gin.H{
	// 		"binding": "error",
	// 		"func":    "create user",
	// 	}
	// 	c.JSON(http.StatusBadRequest, res)
	// 	c.Abort()
	// 	return
	// } else
	if reqBody.First_Name == "" || reqBody.Email == "" || reqBody.Password == "" || reqBody.Phone_no == "" || reqBody.Shop_Name == "" {

		res := gin.H{
			"1 status":           "this field can't be empty",
			"reqBody.First_Name": reqBody.First_Name,
			"reqBody.mail":       reqBody.Email,
			"reqBody.Password":   reqBody.Password,
			"reqBody.Phone_no":   reqBody.Phone_no,
			"reqBody.shop_name":  reqBody.Shop_Name,
			"message":            "if /signupUser/vendor shop name required else text blank in shop_name",
		}
		c.JSON(http.StatusOK, res)

		c.Abort()
		return
	}

	result, hence := pkg.InsertUserDB(reqBody, user_type)

	if !hence {
		resA := gin.H{
			"status": "Account not created",
			"result": result,
		}
		c.JSON(http.StatusBadRequest, resA)
	} else {
		resA := gin.H{
			"status": "Account Succesfully creates",
			"result": result,
		}
		c.JSON(http.StatusOK, resA)
	}

}
