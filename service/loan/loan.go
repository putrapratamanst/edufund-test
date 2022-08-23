package loan

import (
	"edufund-test/interface/loan"
	model "edufund-test/model/loan"
	"edufund-test/presenter"
	"encoding/json"
	"net/http"
	"time"
)

// Service interface
type Service struct {
	repo loan.ILoanRepository
}

// NewService create new use case
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

func (s *Service)Detail(id string) (model.Read, *presenter.Response){
	var result model.Read
	data, errData := s.repo.Get(id)
	if errData != nil {
		return model.Read{}, &presenter.Response{
			Code:    errData.Code,
			Message: errData.Message,
		}
	}

	errJsonUnmarshal := json.Unmarshal([]byte(data), &result)
	if errJsonUnmarshal != nil {
		return model.Read{}, &presenter.Response{
			Code:    http.StatusInternalServerError,
			Message: errJsonUnmarshal.Error(),
		}
	}

	return result, nil
}

func (s *Service)Approve(id string, input model.Read){
	input.IsApproved = true
	input.UpdatedDate = time.Now()
	s.repo.Approve(id, input)
}
