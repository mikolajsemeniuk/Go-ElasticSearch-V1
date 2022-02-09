package application

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/controllers"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/inputs"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/middlewares"
	"github.com/mikolajsemeniuk/go-react-elasticsearch/settings"
)

const (
	postId = ":id"
)

var router = gin.Default()

func Listen() {
	v1 := router.Group(settings.Configuration.GetString("server.basepath"))
	{
		posts := v1.Group("posts")
		{
			posts.GET("", controllers.PostController.All)
			posts.POST("", controllers.PostController.Add)
			posts.GET(postId, controllers.PostController.Single)
			posts.PATCH(postId, middlewares.UUID(postId), middlewares.Body(inputs.Post{}), controllers.PostController.Update)
			posts.DELETE(postId, middlewares.UUID(postId), controllers.PostController.Remove)
		}
	}
	router.Run(fmt.Sprintf(":%s", settings.Configuration.GetString("server.port")))
}
