package main

import (
	"github.com/watariRyo/go-sample-framework/controllers"
	"github.com/watariRyo/go-sample-framework/framework"
)

func main() {
	engine := framework.NewEngine()
	router := engine.Router

	router.Get("/list", controllers.ListController)
	router.Get("/lists/:list_id", controllers.ListItemController)
	router.Get("/lists/:list_id/pictures/:picture_id", controllers.ListItemPictureItemController)
	router.Get("/users", controllers.UsersController)
	router.Get("/students", controllers.StudentsController)

	router.Get("/posts", controllers.PostsPageController)
	router.Post("/posts", controllers.PostsController)
	router.Post("/userposts", controllers.UserPostsController)

	engine.Run()
}
