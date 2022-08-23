package router

import (
	"edufund-test/service/loan"
	controller "edufund-test/controller/loan"

	"github.com/gin-gonic/gin"
)

func Route(v1 *gin.RouterGroup, ln *loan.Service) {
	handler := controller.ControllerHandler(ln)
	loan := v1.Group("/loan")
	{
		loan.POST("", handler.Create)
	}
}