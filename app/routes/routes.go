package routes

import (
	controllers "github.com/danishjsheikh/go-backend-demo/app/controllers"
	"github.com/danishjsheikh/go-backend-demo/app/middleware"
	_ "github.com/danishjsheikh/go-backend-demo/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init(app *gin.Engine) {
	v1 := app.Group("v1")
	{
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		appUser := v1.Group("/users")
		{
			userController := new(controllers.UserController)
			appUser.GET("/", userController.GetUsers)
			appUser.POST("/", userController.CreateUser)
			appUser.PUT("/:username", userController.UpdateUser)
			appUser.DELETE("/:username", userController.DeleteUser)
		}

		appOrder := v1.Group("/orders")
		{
			appOrder.Use(middleware.AddCorrelationID())
			orderController := new(controllers.OrderController)
			appOrder.GET("/:orderCode", orderController.GetOrderByCode)
			appOrder.POST("/", orderController.CreateOrder)
			// appOrder.PUT("/:orderCode", orderController.UpdateOrder)
			// appOrder.DELETE("/:orderCode", orderController.DeleteOrder)
		}
	}
}
