package routes

import (
	"blog/api/controller"
	"blog/infrastructure"
)

//PostRoute -> Route for question module
type ImageRoute struct {
	Controller controller.ImageController
	Handler    infrastructure.GinRouter
}

//NewImageRoute -> initializes new choice routes
func NewImageRoute(
	controller controller.ImageController,
	handler infrastructure.GinRouter,

) ImageRoute {
	return ImageRoute{
		Controller: controller,
		Handler:    handler,
	}
}

//Setup -> setups new choice Routes
func (p ImageRoute) Setup() {
	post := p.Handler.Gin.Group("/posts") //Router group
	{
		post.GET("/images/", p.Controller.GetImages)
		post.POST("/images/", p.Controller.AddImage)
		post.GET("/images/:id", p.Controller.GetImage)
		post.DELETE("/:id", p.Controller.DeleteImage)
		post.PUT("/images/:id", p.Controller.UpdateImage)
	}
}
