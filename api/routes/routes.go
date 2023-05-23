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
        authRoutes.POST("/sign-up", handlers.SignUp)
        authRoutes.POST("/sign-in", handlers.SignInWithEmail)
        authRoutes.POST("/sign-in-jwt", handlers.SignInWithJwt)
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

    // Policies routes
    policiesRoutes := router.Group("/policies")
    policiesRoutes.Use(middleware.AuthRequired())
    {
        policiesRoutes.GET("", handlers.GetAllPolicies)
        policiesRoutes.GET("/:id", handlers.GetPolicyByID)
        policiesRoutes.POST("", handlers.CreatePolicy)
        policiesRoutes.PUT("/:id", handlers.UpdatePolicy)
        policiesRoutes.DELETE("/:id", handlers.DeletePolicy)

        policiesRulesRoutes := policiesRoutes.Group("/:id/rules")
        {
            policiesRulesRoutes.GET("", handlers.GetPolicyRules)
            policiesRulesRoutes.POST("", handlers.AddPolicyRule)
            policiesRulesRoutes.PUT("/:ruleId", handlers.UpdatePolicyRule)
            policiesRulesRoutes.DELETE("/:ruleId", handlers.DeletePolicyRule)
        }
    }

    // Resources routes
    resourcesRoutes := router.Group("/resources")
    resourcesRoutes.Use(middleware.AuthRequired())
    {
        resourcesRoutes.GET("", handlers.GetAllResources)
        resourcesRoutes.GET("/:id", handlers.GetResourceByID)
        resourcesRoutes.POST("", handlers.CreateResource)
        resourcesRoutes.PUT("/:id", handlers.UpdateResource)
        resourcesRoutes.DELETE("/:id", handlers.DeleteResource)

        resourcesAttributesRoutes := resourcesRoutes.Group("/:id/attributes")
        {
            resourcesAttributesRoutes.GET("", handlers.GetResourceAttributes)
            resourcesAttributesRoutes.POST("", handlers.AddResourceAttribute)
            resourcesAttributesRoutes.PUT("/:attributeId", handlers.UpdateResourceAttribute)
            resourcesAttributesRoutes.DELETE("/:attributeId", handlers.DeleteResourceAttribute)
        }
    }

    // Roles routes
    rolesRoutes := router.Group("/roles")
    rolesRoutes.Use(middleware.AuthRequired())
    {
        rolesRoutes.GET("", handlers.GetAllRoles)
        rolesRoutes.GET("/:id", handlers.GetRoleByID)
        rolesRoutes.POST("", handlers.CreateRole)
        rolesRoutes.PUT("/:id", handlers.UpdateRole)
        rolesRoutes.DELETE("/:id", handlers.DeleteRole)

        rolesAttributesRoutes := rolesRoutes.Group("/:id/attributes")
        {
            rolesAttributesRoutes.GET("", handlers.GetRoleAttributes)
            rolesAttributesRoutes.POST("", handlers.AddRoleAttribute)
            rolesAttributesRoutes.PUT("/:attributeId", handlers.UpdateRoleAttribute)
            rolesAttributesRoutes.DELETE("/:attributeId", handlers.DeleteRoleAttribute)
        }
    }

    // Users routes
    usersRoutes := router.Group("/users")
    usersRoutes.Use(middleware.AuthRequired())
    {
        usersRoutes.GET("", handlers.GetAllUsers)
        usersRoutes.GET("/:id", handlers.GetUserByID)
        usersRoutes.POST("", handlers.CreateUser)
        usersRoutes.PUT("/:id", handlers.UpdateUser)
        usersRoutes.DELETE("/:id", handlers.DeleteUser)

        usersRoutes.POST("/:id/roles", handlers.AssignRoleToUser)
        usersRoutes.POST("/:id/resources", handlers.AssignResourceToUser)
        usersRoutes.POST("/:id/policies", handlers.AssignPolicyToUser)

        usersRoutes.DELETE("/:id/roles/:roleId", handlers.RemoveRoleFromUser)
        usersRoutes.DELETE("/:id/resources/:resourceId", handlers.RemoveResourceFromUser)
        usersRoutes.DELETE("/:id/policies/:policyId", handlers.RemovePolicyFromUser)
    }
}
