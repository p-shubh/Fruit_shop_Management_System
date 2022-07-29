package main

import (
	"fruit_shop_management_system/CMD/servers"
	"fruit_shop_management_system/CMD/utils"
)

func main() {

	// DB CONNECTION:-------------------------------

	utils.Connection_with_db()
	defer utils.DB.Close() //--------------------DB get CLOSED

	// REDIS CONNECTION
	utils.Redis()

	// API SERVER
	servers.Servers()

}
