package api

import (
	"github.com/Imomali1/gophermart/internal/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.POST("/api/user/register", Register)
	router.POST("/api/user/login", Login)

	authorized := router.Group("/api/user")
	authorized.Use(middlewares.IsAuthorized())
	{
		authorized.POST("/orders", LoadOrders)
		authorized.GET("/orders", ListOrders)
		authorized.GET("/balance", GetBalance)
		authorized.POST("/balance/withdraw", Withdraw)
		authorized.GET("/withdrawals", GetWithdrawalInfo)
	}

	return router
}
