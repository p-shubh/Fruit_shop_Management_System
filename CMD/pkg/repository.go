package pkg

import (
	"fmt"
	"fruit_shop_management_system/CMD/utils"
)

func InsertUserDB(reqBody Create_Shop_Account, user_type int) (Create_Shop_Account, bool) {

	var hence bool = true
	// var a = reqBody

	// fmt.Println("a", a)

	if user_type == 2 {
		reqBody.Shop_Name = "general user account"
	}

	Insert := "INSERT INTO account (first_name,last_name,address,email,shop_name,password,user_type) VALUES ($1,$2,$3,$4,$5 ,$6,$7);"

	_, err := utils.DB.Exec(Insert, reqBody.First_Name, reqBody.Last_Name, reqBody.Address, reqBody.Email, reqBody.Shop_Name, reqBody.Password, user_type)

	if err != nil { //there is an error

		hence = false

		fmt.Println(err)

		panic(err)

	} else {
		hence = true
	}

	return reqBody, hence

	// return reqBody
}
