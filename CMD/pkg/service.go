package pkg

import (
	"encoding/json"
	"fmt"
	"fruit_shop_management_system/CMD/utils"
	"time"
)

func InsertFruitService(reqBody Add_Fruits) (Add_Fruits, bool) {

	res := InsertFruitsDB(reqBody)

	bools := true

	if res {

		data, bools := GetAllOrders()

		if !bools {
			fmt.Println("values", data)
			fmt.Println("error fail to GetAllOrders() in InsertFruitService/service.go")
		}

		jsonData, _ := json.Marshal(data)

		r := utils.Client.Set("all", jsonData, time.Minute*6)

		fmt.Println("r := test ", r)

	} else if !res {
		fmt.Println("values", res)
		fmt.Println("error fail to InsertFruitsDB in InsertFruitService/service.go")
	}
	return reqBody, bools

}

func GetAllOrdersService() ([]Add_Fruits, bool) {

	data, bools := GetAllOrders()

	if !bools {
		fmt.Println("values", data)
		fmt.Println("error fail to GetAllOrders() in InsertFruitService/service.go")
	}

	jsonData, _ := json.Marshal(data)

	r2 := utils.Client.Set("all", jsonData, time.Minute*6)

	fmt.Println("r2 := test ", r2)

	return data, bools

}
