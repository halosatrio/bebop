package main

import (
	"github.com/halosatrio/bebop/config"
	"github.com/halosatrio/bebop/router"
)

func main() {
	config.LoadEnv()

	db := config.ConnectDB()
	defer db.Close()

	r := router.SetupRouter(db)
	r.Run(":8080")
}
