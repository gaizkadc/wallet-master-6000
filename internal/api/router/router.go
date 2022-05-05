package router

import (
	"github.com/gaizkadc/wallet-master-6000/internal/api/controllers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	app := gin.New()

	app.GET("/health", controllers.HealthCheck)
	app.GET("/customers/:id", controllers.GetCustomerById)
	app.GET("/customers/:id/transactions", controllers.GetTransactionsByCustomerId)
	app.POST("/transactions", controllers.AddTransaction)

	return app
}
