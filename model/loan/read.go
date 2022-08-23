package loan

import "time"

type Read struct {
	Price       string    `json:"price"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
	Period      string    `json:"period"`
	IsApproved  bool      `json:"isApproved"`
}
