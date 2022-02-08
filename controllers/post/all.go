package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*postController) All(context *gin.Context) {
	context.JSON(http.StatusOK, "All")
}
