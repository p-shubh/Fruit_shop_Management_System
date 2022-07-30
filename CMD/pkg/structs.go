package pkg

type Create_Shop_Account struct {
	Create_User
	Shop_Name string `json:"shop_name" binding:"required"`
}
type Create_User struct {
	First_Name string `json:"first_name" binding:"required"`
	Last_Name  string `json:"last_name"`
	Address    string `json:"address"`
	Phone_no   string `json:"phone_no" binding:"required"`
	Login             //"email","password"
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,alphanum"`
}

type Add_Fruits struct {
	Id int `json:"id"`
	Fruits string `json:"Fruits" binding:"required"`
}

type Id struct {
	Id int `json:"id" binding:"required"`
}

type Get_User_Detail struct {
	Create_Shop_Account
	Id        int `json:"id"`
	User_Type int `json:"user_type"`
}
