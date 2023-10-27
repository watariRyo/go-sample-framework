package controllers

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"

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
	name := ctx.FormKey("name", "defaultName")
	age := ctx.FormKey("age", "30")
	fileInfo, err := ctx.FormFile("file")
	if err != nil {
		ctx.WriteHeader(http.StatusInternalServerError)
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s %s %s", name, age, fileInfo.Filename), fileInfo.Data, fs.ModePerm)
	if err != nil {
		ctx.WriteHeader(http.StatusInternalServerError)
	}

	ctx.WriteString("success")
}

func PostsPageController(ctx *framework.MyContext) {
	ctx.WriteString(`<!DOCTYPE html>
	<html>
		<head>
			<title>form</title>
		</head>
		<body>
			<div>
				<form action="/posts" method="post" enctype="multipart/form-data">
					<div><label>name</label>: <input name="name"/></div>
					<div><label>age</label>: 
					<select name="age">
						<option value="1">1</option>
						<option value="2">2</option>
					</select></div>
					<button type="submit">submit</button>
					<input name="file" type="file"/>
				</form>
			</div>
		</body>
	</html>`)
}
