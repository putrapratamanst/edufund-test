package loan

import "time"

type Update struct {
	Price       string    `json:"price" validate:"required"`
	Period      string    `json:"period" validate:"required"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
	IsApproved  bool      `json:"isApproved"`
}
