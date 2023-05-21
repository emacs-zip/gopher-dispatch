package routes

import (
	"gopher-dispatch/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
   // Authentication routes
   authRoutes := router.Group("/auth")
   {
       authRoutes.POST("/sign-in", handlers.SignInWithEmail)
       authRoutes.POST("/sign-in-jwt", handlers.SignInWithJwt)
       authRoutes.POST("/sign-up", handlers.SignUp)
       authRoutes.GET("/forgot-password", handlers.ForgotPassword)
   }


   // Analytics routes
   analyticsRoutes := router.Group("/analytics")
   {
       analyticsRoutes.POST("/page-view", handlers.RecordPageView)
       analyticsRoutes.GET("/page-view/:user_id", handlers.GetUserPageView)
   }
}
