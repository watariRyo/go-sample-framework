package controllers

import (
	"github.com/watariRyo/go-sample-framework/framework"
)

type StudentResponse struct {
	Name string `json:"name"`
}

func StudentsController(ctx *framework.MyContext) {
	name := ctx.QueryKey("name", "")

	studentResponse := &StudentResponse{
		Name: name,
	}

	ctx.Json(studentResponse)
	return
}

func ListController(ctx *framework.MyContext) {
	ctx.WriteString("list")
}

func ListItemController(ctx *framework.MyContext) {
	ctx.WriteString("listItem")
}

func ListItemPictureItemController(ctx *framework.MyContext) {
	listID := ctx.GetParam(":list_id", "")
	pictureID := ctx.GetParam(":picture_id", "")

	type OUTPUT struct {
		ListID    string `json:"list_id"`
		PictureID string `json:"picture_id"`
	}

	output := &OUTPUT{
		ListID:    listID,
		PictureID: pictureID,
	}

	ctx.Json(output)
}

func UsersController(ctx *framework.MyContext) {
	ctx.WriteString("users")
}

func PostsController(ctx *framework.MyContext) {
	ctx.WriteString("posts")
}

func PostsPageController(ctx *framework.MyContext) {
	ctx.WriteString(`<!DOCTYPE html>
	<html>
		<head>
			<title>form</title>
		</head>
		<body>
			<div>
				<form action="/posts" method="post">
					<input name="name"/>
					<button type="submit">submit</submit>
				</form>
			</div>
		</body>
	</html>`)
}
