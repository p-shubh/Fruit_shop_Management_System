package pkg

type Create_User struct {
	First_Name string `json:"first_name" binding:"required"`
	Last_Name  string `json:"last_name"`
	Address    string `json:"address"`
	Phone_no   string `json:"phone_no" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required,alphanum"`
}

type Create_Shop_Account struct {
	Create_User
	Shop_Name string `json:"shop_name" binding:"required"`
}

type Add_Fruits struct {
	Id     int    `json:"id" binding:"required"`
	Fruits string `json:"Fruits" binding:"required"`
}

type Id struct {
	Id int `json:"id"`
}
