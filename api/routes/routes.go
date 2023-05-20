package routes

import (
	"gopher-dispatch/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
   // Authentication routes
   r.POST("/auth/sign-in", handlers.SignInWithEmail)
   r.POST("/auth/sign-in-jwt", handlers.SignInWithJwt)
   r.POST("/auth/sign-up", handlers.SignUp)
   r.GET("/auth/forgot-password", handlers.ForgotPassword)

   // Analytics routes
   r.POST("/analytics/page-view", handlers.RecordPageView)
}
