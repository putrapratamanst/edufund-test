package loan

import (
	"edufund-test/model/loan"
	"edufund-test/pkg"
	"edufund-test/presenter"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Controller) Update(ctx *gin.Context) {
	var input loan.Update
	id := ctx.Param("id")
	err := pkg.ValidateRequest(pkg.BIND_TYPE_JSON, "Loan Create", ctx, &input)
	if err != nil {
		pkg.Response(ctx, &presenter.Response{
			Code:    err.Code,
			Message: err.Message,
		})
		return
	}

	data, errData := handler.ln.Detail(id)
	if errData != nil {
		pkg.Response(ctx, &presenter.Response{
			Code:    errData.Code,
			Message: errData.Message,
		})
		return
	}

	input.UpdatedDate = data.UpdatedDate
	input.CreatedDate = data.CreatedDate
	input.IsApproved = data.IsApproved
	
	handler.ln.Change(id, input)
	pkg.Response(ctx, &presenter.Response{
		Code: http.StatusOK,
		Message: "Successfully Update Loan",
	})
}