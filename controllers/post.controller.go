package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/inputs"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/services"
)

var PostController IPostControler = postController{}

type IPostControler interface {
	All(context *gin.Context)
	Add(context *gin.Context)
	Single(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
}

type postController struct{}

func (postController) All(context *gin.Context) {
	posts, err := services.PostService.All()
	if err != nil {
		context.JSON(http.StatusServiceUnavailable, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": posts})
}

func (postController) Add(context *gin.Context) {
	input := inputs.Post{
		Title: "my title",
		Done:  false,
	}

	err := services.PostService.Add(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	context.JSON(http.StatusOK, "Add")
}

func (postController) Single(context *gin.Context) {
	context.JSON(http.StatusOK, "Single")
}

func (postController) Update(context *gin.Context) {
	context.JSON(http.StatusOK, "Update")
}

func (postController) Remove(context *gin.Context) {
	context.JSON(http.StatusOK, "Remove")
}
