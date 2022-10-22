package router

import (
	"final/controllers"
	"final/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:userId", middlewares.Authentication(), controllers.UpdateUsers)
		userRouter.DELETE("/:userId", middlewares.Authentication(), controllers.DeleteUsers)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.ShowPhoto)
		photoRouter.PUT("/:photoId", controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", controllers.ShowComment)
		commentRouter.PUT("/:commentId", controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", controllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.ShowSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", controllers.DeleteSocialMedia)
	}
	return r
}
