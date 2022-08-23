package loan

import (
	"edufund-test/pkg"
	"edufund-test/presenter"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Controller) Approve(ctx *gin.Context) {
	id := ctx.Param("id")
	data, errData := handler.ln.Detail(id)
	if errData != nil {
		pkg.Response(ctx, &presenter.Response{
			Code:    errData.Code,
			Message: errData.Message,
		})
		return
	}

	handler.ln.Approve(id, data)
	pkg.Response(ctx, &presenter.Response{
		Code:    http.StatusOK,
		Message: "Successfully Approve Loan",
	})
}