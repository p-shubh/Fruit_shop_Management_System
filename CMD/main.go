package main

import (
	"fruit_shop_management_system/CMD/servers"
	"fruit_shop_management_system/CMD/utils"
)

func main() {

	utils.Connection_with_db()
	defer utils.DB.Close()

	servers.Servers()
}
