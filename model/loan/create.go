package loan

import "time"

type Create struct {
	Price       string    `json:"price" validate:"required"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
	Period      string    `json:"period" validate:"required"`
	IsApproved  bool      `json:"isApproved"`
}
