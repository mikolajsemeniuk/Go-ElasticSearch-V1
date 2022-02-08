package application

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/controllers"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/settings"
)

var router = gin.Default()

func Listen() {
	router.GET("", controllers.PostController.All)
	router.POST("", controllers.PostController.Add)
	router.GET(":id", controllers.PostController.Single)
	router.PATCH(":id", controllers.PostController.Update)
	router.DELETE(":id", controllers.PostController.Remove)
	router.Run(fmt.Sprintf(":%s", settings.Configuration.GetString("server.port")))
}
