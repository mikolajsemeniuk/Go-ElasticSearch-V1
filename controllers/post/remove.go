package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*postController) Remove(context *gin.Context) {
	context.JSON(http.StatusOK, "Remove")
}
