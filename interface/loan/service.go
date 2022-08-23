package loan

import "edufund-test/model/loan"

type ILoanService interface {
	Create(input loan.Create)
	// Read()
	// Update()
	// Approve()
}