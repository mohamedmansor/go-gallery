package main

import (
	"blog/api/controller"
	"blog/api/repository"
	"blog/api/routes"
	"blog/api/service"
	"blog/infrastructure"
	"blog/models"
)

func init() {
	infrastructure.LoadEnv()
}
func main() {
	router := infrastructure.NewGinRouter()                     //router has been initialized and configured
	db := infrastructure.NewDatabase()                          // database has been initialized and configured
	postRepository := repository.NewPostRepository(db)          // repository are being setup
	postService := service.NewPostService(postRepository)       // service are being setup
	postController := controller.NewPostController(postService) // controller are being set up
	postRoute := routes.NewPostRoute(postController, router)    // post routes are initialized
	postRoute.Setup()                                           // post routes are being setup

	// Image Routes
	imageRepository := repository.NewImageRepository(db)
	imageService := service.NewImageService(imageRepository)
	imageController := controller.NewImageController(imageService)
	imageRoute := routes.NewImageRoute(imageController, router)
	imageRoute.Setup()

	db.DB.AutoMigrate(&models.Post{}) // migrating Post model to datbase table
	router.Gin.Run(":8000")           //server started on 8000 port
}
