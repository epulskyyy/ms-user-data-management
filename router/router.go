package router

import (
	"github.com/gin-gonic/gin"
	"ms-user-data-management/config"
	"ms-user-data-management/controllers"
	"ms-user-data-management/middlewares"
	"ms-user-data-management/repo"
	"ms-user-data-management/services"
	"os"
)

func NewRoute()  {
	router := gin.New()
	port := os.Getenv("PORT")
	db := config.ConnectToDB()
	router.Use(middlewares.CorsMiddleware())
	router.Use(gin.Logger())
	defer db.Close()
	//user
	userRepo := repo.CreateUserRepoImpl(db)
	userService := services.CreateUserServiceImpl(userRepo)
	//pendidikan
	pendidikanRepo := repo.CreatePendidikanRepoImpl(db)
	pendidikanService := services.CreatePendidikanServiceImpl(pendidikanRepo)
	//profesi
	profesiRepo := repo.CreateProfesiRepoImpl(db)
	profesiService := services.CreateProfesiServiceImpl(profesiRepo)
	v1 := router.Group("api/v1")

	{
		controllers.CreateUserController(v1, userService)
		controllers.CreatePendidikanController(v1, pendidikanService)
		controllers.CreateProfesiController(v1, profesiService)
	}

	router.Run(port)
}