package loan

import "edufund-test/infrastructure"

type Repository struct {
	rc *infrastructure.RedisCache
}

// Loan new repository
func NewRepository(rc *infrastructure.RedisCache) *Repository {
	return &Repository{
		rc: rc,
	}
}

func (repo *Repository)Create(){

}
