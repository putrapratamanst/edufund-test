package loan

import (
	"edufund-test/interface/loan"
	model "edufund-test/model/loan"
	"time"
)

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

func (s *Service)Create(input model.Create){
	input.CreatedDate = time.Now()
	input.UpdatedDate = time.Now()
	s.repo.Insert(input)
}