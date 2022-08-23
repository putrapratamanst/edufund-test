package loan

import (
	"edufund-test/model/loan"
	"edufund-test/presenter"
)

type ILoanService interface {
	Create(input loan.Create)
	Detail(id string) (loan.Read, presenter.Response)
	// Update()
	// Approve()
}