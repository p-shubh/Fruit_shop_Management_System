package pkg

import (
	"fmt"
	"fruit_shop_management_system/CMD/utils"

	"github.com/lib/pq"
)

func InsertUserDB(reqBody Create_Shop_Account, user_type int) (Create_Shop_Account, bool) {

	var hence bool = true

	if user_type == 2 {
		reqBody.Shop_Name = "general user account"
	}

	Insert := "INSERT INTO account (first_name,last_name,address,email,shop_name,password,user_type) VALUES ($1,$2,$3,$4,$5 ,$6,$7);"

	_, err := utils.DB.Exec(Insert, reqBody.First_Name, reqBody.Last_Name, reqBody.Address, reqBody.Email, reqBody.Shop_Name, reqBody.Password, user_type)

	if err != nil {

		err2 := UniqueViolation(err)

		if err2 != nil {
			hence = false

			fmt.Println(err)

			panic(err)
		}

		hence = false

		fmt.Println(err)

		panic(err)

	} else {
		hence = true
	}

	return reqBody, hence

}

func InsertFruitsDB(reqBody Add_Fruits) (Add_Fruits, bool) {

	var hence bool = true

	Insert := `INSERT INTO fruits_account ("id", "fruits") VALUES ($1,$2);`

	_, err := utils.DB.Exec(Insert, reqBody.Id, reqBody.Fruits)

	if err != nil {

		hence = false

		fmt.Println(err)

		panic(err)

	} else {
		hence = true
	}

	return reqBody, hence

}

func UpdateFruits(reqBody Add_Fruits) (Add_Fruits, bool) {

	var hence bool = true

	Insert := `update fruits_account set fruits = $1 where id = $2;`

	_, err := utils.DB.Exec(Insert, reqBody.Fruits, reqBody.Id)

	if err != nil {
		hence = false

		fmt.Println(err)

		panic(err)

	} else {
		hence = true
	}

	return reqBody, hence

}

func DeleteFruitItems(Id int) (int, bool) {

	var hence bool = true

	Insert := `delete from fruits_account where id = $1;`

	_, err := utils.DB.Exec(Insert, Id)

	if err != nil {

		hence = false

		fmt.Println(err)

		panic(err)

	} else {
		hence = true
	}

	return Id, hence

}

func FetchOrders(Id int) (Add_Fruits, bool) {

	var hence bool = true

	reqBody := Add_Fruits{}

	row := utils.DB.QueryRow(`SELECT id, fruits FROM fruits_account where id = $1;`, Id)

	rows := row.Scan(&reqBody.Id, &reqBody.Fruits)

	if rows != nil {
		hence = false
	}

	return reqBody, hence

}

func UniqueViolation(err error) *pq.Error {
	if pqerr, ok := err.(*pq.Error); ok &&
		pqerr.Code == "23505" {
		return pqerr
	}
	return nil
}

// ====================================================count user exist=========================================

func EmailExist(Email string) bool {

	// Err := `select count(*) from account where email = $1`

	var count int
	var hence bool = true

	// reqBody := Add_Fruits{}

	row := utils.DB.QueryRow(`select count(*) from account where email = $1;`, Email)

	rows := row.Scan(&count)

	if rows != nil || count == 0 {
		hence = false
		fmt.Println(Email, "doesn't exist")
	}

	return hence
}

// =====================================================get user===================================================

func GetUserByEmail(Email string)  {
	
}
