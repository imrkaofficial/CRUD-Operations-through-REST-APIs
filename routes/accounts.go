package routes

import (
	"mysql-api/controller"

	"github.com/gin-gonic/gin"
)

func AccRoute(router *gin.Engine) {
	account := router.Group("/Account")
	{
		account.GET("", controller.GetAcc)
		account.POST("", controller.CreateAcc)
		account.PUT("/:uid", controller.UpdateAcc)
		account.DELETE("/:uid", controller.DelAcc)
	}
}
