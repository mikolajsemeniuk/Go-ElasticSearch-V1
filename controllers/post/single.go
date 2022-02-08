package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*postController) Single(context *gin.Context) {
	context.JSON(http.StatusOK, "Single")
}
