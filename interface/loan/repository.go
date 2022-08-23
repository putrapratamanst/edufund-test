package loan

import (
	model "edufund-test/model/loan"
	"edufund-test/presenter"
)

type ILoanRepository interface {
	Insert(input model.Create)
	Get(id string) (string, *presenter.Response)
	// Update()
	// Approve()
}
