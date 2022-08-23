package loan

import (
	model "edufund-test/model/loan"
)

type ILoanRepository interface {
	Insert(input model.Create)
	// Read()
	// Update()
	// Approve()
}
