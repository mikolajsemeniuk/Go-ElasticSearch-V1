package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*postController) Add(context *gin.Context) {
	context.JSON(http.StatusOK, "Add")
}
