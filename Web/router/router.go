package router

import (
	"github.com/gin-gonic/gin"
	services "react-go/Controllers"
	"react-go/Web/middlewares"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	v1 := r.Group("/")
	{
		// Login Route
		v1.POST("login", services.Login)

		//// Reset password:
		//v1.POST("/password/forgot", services.ForgotPassword)
		//v1.POST("/password/reset", services.ResetPassword)

		//Users routes
		v1.POST("users", services.CreateUser)
		v1.GET("users", services.GetUsers)
		v1.GET("users/:id", services.GetUser)
		v1.PUT("users/:id", middlewares.TokenAuthMiddleware(), services.UpdateUser)
		v1.PUT("avatar/users/:id", middlewares.TokenAuthMiddleware(), services.UpdateAvatar)
		v1.DELETE("users/:id", middlewares.TokenAuthMiddleware(), services.DeleteUser)

		//Posts routes
		v1.POST("/posts", middlewares.TokenAuthMiddleware(), services.CreatePost)
		v1.GET("/posts", services.GetPosts)
		v1.GET("/posts/:id", services.GetPost)
		v1.PUT("/posts/:id", middlewares.TokenAuthMiddleware(), services.UpdatePost)
		v1.DELETE("/posts/:id", middlewares.TokenAuthMiddleware(), services.DeletePost)
		v1.GET("/user_posts/:id", services.GetUserPosts)

		//Like route
		v1.GET("/likes/:id", services.GetLikes)
		v1.POST("/likes/:id", middlewares.TokenAuthMiddleware(), services.LikePost)
		v1.DELETE("/likes/:id", middlewares.TokenAuthMiddleware(), services.UnLikePost)

		//Comment routes
		v1.POST("/comments/:id", middlewares.TokenAuthMiddleware(), services.CreateComment)
		v1.GET("/comments/:id", services.GetComments)
		v1.PUT("/comments/:id", middlewares.TokenAuthMiddleware(), services.UpdateComment)
		v1.DELETE("/comments/:id", middlewares.TokenAuthMiddleware(), services.DeleteComment)
	}

	return r
}

