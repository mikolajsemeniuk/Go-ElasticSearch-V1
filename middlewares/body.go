package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/inputs"
)

func Body(input interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {

		switch input.(type) {
		case inputs.Post:
			input = &inputs.Post{}
		}

		if err := context.BindJSON(&input); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}

		context.Set("input", input)
	}
}
