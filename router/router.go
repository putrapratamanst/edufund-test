package router

import (
	controller "edufund-test/controller/loan"
	"edufund-test/service/loan"

	"github.com/gin-gonic/gin"
)

func Route(v1 *gin.RouterGroup, ln *loan.Service) {
	handler := controller.ControllerHandler(ln)
	loan := v1.Group("/loan")
	{
		loan.POST("", handler.Create)
		loan.GET("/:id", handler.Read)
		loan.PUT("/:id", handler.Approve)
	}
}