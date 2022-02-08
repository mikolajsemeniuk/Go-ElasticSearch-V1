package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (*postController) Update(context *gin.Context) {
	context.JSON(http.StatusOK, "Update")
}
