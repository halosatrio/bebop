package main

import (
	"github.com/halosatrio/bebop/config"
	"github.com/halosatrio/bebop/db"
	"github.com/halosatrio/bebop/router"
)

func main() {
	config.LoadEnv()
	r := router.SetupRouter()

	db, _ := db.InitDB()
	defer db.Close()

	r.Run(":8080")
}
