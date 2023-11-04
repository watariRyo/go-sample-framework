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

	router.Get("/json_p_test", controllers.JsonPTestController)

	router.Use(framework.AuthUserMiddleware)
	router.Use(framework.TimeCostMiddleware)
	router.Use(framework.TimeOutMiddleWare)

	router.UseNoRoute(func(ctx *framework.MyContext) {
		ctx.WriteString("not found")
	})

	router.Use(framework.StaticFileMiddleware)

	engine.Run()
}
