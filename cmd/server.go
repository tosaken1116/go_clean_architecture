package main

import (
	"go-clean-archtecture/api/adapter/http/router"
	"go-clean-archtecture/api/infra"
)

func main() {
	infra.MigrateDatabase()
	r := router.InitRouter()
	r.Run()
}
