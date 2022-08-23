package loan

import (
	"edufund-test/model/loan"
	"edufund-test/pkg"
	"edufund-test/presenter"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Controller) Create(ctx *gin.Context) {
	var input loan.Create
	err := pkg.ValidateRequest(pkg.BIND_TYPE_JSON, "Loan Create", ctx, &input)
	if err != nil {
		pkg.Response(ctx, &presenter.Response{
			Code:    err.Code,
			Message: err.Message,
		})
		return
	}

	handler.ln.Create(input)
	pkg.Response(ctx, &presenter.Response{
		Code: http.StatusOK,
		Message: "Successfully Create Loan",
	})
}