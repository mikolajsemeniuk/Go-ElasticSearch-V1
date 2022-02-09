package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UUID(routeName string) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, err := uuid.Parse(context.Param("id"))

		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"errors": fmt.Sprintf("%s is not valid, %s", routeName, err.Error()),
				},
			)
			return
		}

		context.Set("id", id)
	}
}
