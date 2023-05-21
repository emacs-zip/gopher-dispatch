package routes

import (
	"gopher-dispatch/api/handlers"
	"gopher-dispatch/api/middleware"

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
   analyticsRoutes.Use(middleware.AuthRequired())
   {
       analyticsRoutes.POST("/page-view", handlers.RecordPageView)
   }
   analyticsRoutes.Use(middleware.AuthRequired(), middleware.RBAC("admin"))
   {
       analyticsRoutes.GET("/page-view/:user_id", handlers.GetUserPageView)
   }
}
