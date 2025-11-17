package route

//import (
//	"gau-cloud-orchestrator/controller"
//	"gau-cloud-orchestrator/middleware"
//
//	//"gau-cloud-orchestrator/middleware"
//
//	"github.com/gin-gonic/gin"
//	"github.com/tnqbao/gau-cloud-orchestrator/repository/database"
//	"gorm.io/gorm"
//)
//
//func SetupVMRoutes(r *gin.Engine, ctrl *controller.Controller, jwtMiddleware gin.HandlerFunc) {
//	vmGroup := r.Group("/api/v1/vms")
//	vmGroup.Use(jwtMiddleware) // Apply JWT authentication middleware
//	{
//		vmGroup.POST("", ctrl.CreateVM)
//		vmGroup.GET("", ctrl.ListVMs)
//		vmGroup.GET("/:id", ctrl.GetVM)
//		vmGroup.DELETE("/:id", ctrl.DeleteVM)
//	}
//}
//
//func SetupRoutes(r *gin.Engine, ctrl *controller.Controller) {
//	// Health check
//	r.GET("/health", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"status": "healthy",
//		})
//	})
//
//	// JWT middleware (you need to implement this in middleware package)
//	jwtMiddleware := middleware.JWTAuthMiddleware()
//
//	// Setup VM routes
//	SetupVMRoutes(r, ctrl, jwtMiddleware)
//}
