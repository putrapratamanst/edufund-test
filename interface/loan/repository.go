package loan

import (
	"edufund-test/model/loan"
	model "edufund-test/model/loan"
	"edufund-test/presenter"
)

type ILoanRepository interface {
	Insert(input model.Create)
	Get(id string) (string, *presenter.Response)
	Approve(id string, input loan.Read)
	Change(id string, input loan.Update)
}
