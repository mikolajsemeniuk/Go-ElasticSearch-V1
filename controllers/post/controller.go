package post

import "github.com/gin-gonic/gin"

var PostController IPostControler = &postController{}

type IPostControler interface {
	All(context *gin.Context)
	Add(context *gin.Context)
	Single(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
}

type postController struct{}
