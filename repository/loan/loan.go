package loan

import (
	"context"
	"edufund-test/infrastructure"
	model "edufund-test/model/loan"
	"edufund-test/pkg"
	"edufund-test/presenter"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
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

func (repo *Repository) Get(id string) (string, *presenter.Response) {
	context := context.Background()
	key := fmt.Sprintf("loan:%s", id)
	data, err := repo.rc.Client.Get(context, key).Result()
	if err != nil && err != redis.Nil {
		return "", &presenter.Response{
			Message: pkg.ErrGetDataRedis.Error(),
		}
	}
	return data, nil

}

func (repo *Repository) Approve(id string, input model.Read) {
	context := context.Background()
	encodeData, _ := json.Marshal(input)
	repo.rc.Client.Set(context, "loan:"+id, encodeData, 0)
}

func (repo *Repository) Change(id string, input model.Update) {
	context := context.Background()
	encodeData, _ := json.Marshal(input)
	repo.rc.Client.Set(context, "loan:"+id, encodeData, 0)
}
