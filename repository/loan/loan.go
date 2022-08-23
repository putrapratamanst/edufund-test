package loan

import (
	"context"
	"edufund-test/infrastructure"
	model "edufund-test/model/loan"
	"edufund-test/pkg"
	"encoding/json"
)

type Repository struct {
	rc *infrastructure.RedisCache
}

// Loan new repository
func NewRepository(rc *infrastructure.RedisCache) *Repository {
	return &Repository{
		rc: rc,
	}
}

func (repo *Repository) Insert(input model.Create) {
	id := pkg.RandomString()
	context := context.Background()
	encodeData, _ := json.Marshal(input)
	repo.rc.Client.Set(context, "loan:"+id, encodeData, 0)
}
