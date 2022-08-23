package loan

import "edufund-test/interface/loan"

//Service interface
type Service struct {
	repo loan.ILoanRepository
}

//NewService create new use case
func NewService(repo loan.ILoanRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func Create(){
	
}