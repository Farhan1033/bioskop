package main

import (
	"bioskop/handler"
	"bioskop/infra"
	"bioskop/infra/config"
	bioskoprepositorypg "bioskop/repository/bioskop_repository_pg"
	bioskopserviceimpl "bioskop/service/bioskop_service_impl"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	infra.InitDB()

	r := gin.Default()

	//Init Repository
	bioskopRepo := bioskoprepositorypg.NewBioskopRepository(infra.DB)

	//Init Service
	bioskopService := bioskopserviceimpl.NewBioskopService(bioskopRepo)

	//Setup Router
	api := r.Group("/api/v1")

	//Init Handler
	handler.NewBioskopHandler(api, bioskopService)

	r.Run(":" + config.GetKey("APP_PORT"))
}
