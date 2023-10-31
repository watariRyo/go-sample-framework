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
	list := make([]string, 0)
	ctx.WriteString(list[3])
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

type PostsPageForm struct {
	Name string
}

func PostsPageController(ctx *framework.MyContext) {
	authUser := ctx.Get("AuthUser", "defaultValue")
	postsPageForm := &PostsPageForm{
		Name: authUser.(string),
	}
	ctx.RenderHtml("./html/posts_page.html", postsPageForm)
}

type UserPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func UserPostsController(ctx *framework.MyContext) {
	userPost := &UserPost{}
	if err := ctx.BindJson(userPost); err != nil {
		ctx.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx.Json(userPost)
}

func JsonPTestController(ctx *framework.MyContext) {
	queryKey := ctx.QueryKey("callback", "cb")

	ctx.JsonP(queryKey, "")
}
