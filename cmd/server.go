package main

import (
	"go-clean-architecture/api/adapter/http/router"
	"go-clean-architecture/api/infra"
)

func main() {
	infra.MigrateDatabase()
	r := router.InitRouter()
	r.Run()
}
