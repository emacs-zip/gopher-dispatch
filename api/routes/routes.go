package routes

import (
	"gopher-dispatch/api/handlers"
	"gopher-dispatch/api/middleware"

	"github.com/gin-gonic/gin"
)


func SetupRouter(router *gin.Engine) {
    // Authentication
    authRoutes := router.Group("/auth")
    {
        authRoutes.POST("/sign-up", handlers.SignUp)
        authRoutes.POST("/sign-in", handlers.SignInWithEmail)
        authRoutes.POST("/sign-in-jwt", handlers.SignInWithJwt)
        authRoutes.GET("/forgot-password", handlers.ForgotPassword)
    }

    // Analytics
    analyticsRoutes := router.Group("/analytics")
    analyticsRoutes.Use(middleware.AuthRequired(), middleware.ABAC())
    {
        analyticsRoutes.POST("/page-view", handlers.RecordPageView)
    }
    analyticsRoutes.Use(middleware.AuthRequired(), middleware.ABAC())
    {
        analyticsRoutes.GET("/page-view/:userID", handlers.GetUserPageView)
    }

    // User routes
    userRoutes := router.Group("/users")
    userRoutes.Use(middleware.AuthRequired(), middleware.ABAC())
    {
        userRoutes.GET("", handlers.GetUsers)
        userRoutes.GET("/:id", handlers.GetUser)
        userRoutes.POST("", handlers.CreateUser)
        userRoutes.PUT("/:id", handlers.UpdateUser)
        userRoutes.DELETE("/:id", handlers.DeleteUser)
        userRoutes.GET("/:id/attributes", handlers.GetUserAttributes)
        userRoutes.GET("/:id/attributes/:attributeId", handlers.GetUserAttribute)
        userRoutes.POST("/:id/attributes", handlers.AddUserAttribute)
        userRoutes.PUT("/:id/attributes/:attributeId", handlers.UpdateUserAttribute)
        userRoutes.DELETE("/:id/attributes/:attributeId", handlers.DeleteUserAttribute)
    }

    // Attribute routes
    attributeRoutes := router.Group("/attributes")
    attributeRoutes.Use(middleware.AuthRequired(), middleware.ABAC())
    {
        attributeRoutes.GET("", handlers.GetAttributes)
        attributeRoutes.GET("/:id", handlers.GetAttribute)
        attributeRoutes.POST("", handlers.CreateAttribute)
        attributeRoutes.PUT("/:id", handlers.UpdateAttribute)
        attributeRoutes.DELETE("/:id", handlers.DeleteAttribute)
    }

    // Policy routes
    policyRoutes := router.Group("/policies")
    policyRoutes.Use(middleware.AuthRequired(), middleware.ABAC())
    {
        policyRoutes.GET("", handlers.GetPolicies)
        policyRoutes.GET("/:id", handlers.GetPolicy)
        policyRoutes.POST("", handlers.CreatePolicy)
        policyRoutes.PUT("/:id", handlers.UpdatePolicy)
        policyRoutes.DELETE("/:id", handlers.DeletePolicy)
    }

    // Tenant routes
    tenantRoutes := router.Group("/tenants")
    tenantRoutes.Use(middleware.AuthRequired(), middleware.ABAC())
    {
        tenantRoutes.GET("", handlers.GetTenants)
        tenantRoutes.GET("/:id", handlers.GetTenant)
        tenantRoutes.POST("", handlers.CreateTenant)
        tenantRoutes.PUT("/:id", handlers.UpdateTenant)
        tenantRoutes.DELETE("/:id", handlers.DeleteTenant)
        tenantRoutes.GET("/:id/users", handlers.GetTenantUsers)
        tenantRoutes.GET("/:id/policies", handlers.GetTenantPolicies)
    }
}
