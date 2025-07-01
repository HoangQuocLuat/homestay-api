package handler

import (
	"back-end/cmd/handler/homestay"
	"back-end/cmd/svc"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(router *gin.Engine, serverCtx *svc.ServiceContext) {
	router.LoadHTMLGlob("cmd/templates/pages/*.tmpl")

	//homestay
	homestayGroup := router.Group("/homestay")
	{
		homestayGroup.POST("", homestay.CreateHomestayHandler(serverCtx))
		homestayGroup.GET("/all", homestay.GetAllHomestayHandler(serverCtx))
		homestayGroup.GET("/:id", homestay.GetHomestayHandler(serverCtx))
		homestayGroup.PUT("/:id", homestay.UpdateHomestayHandler(serverCtx))
	}

}
