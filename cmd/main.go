package main

import (
	"edufund-test/infrastructure"
	"edufund-test/pkg"
	"edufund-test/presenter"
	repositoryLoan "edufund-test/repository/loan"
	"edufund-test/router"
	serviceLoan "edufund-test/service/loan"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	
	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")
	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	fmt.Println("Listening LOAN API ...", serverString)

	r := gin.Default()
	redisConfig := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	redisCache, errRedis := infrastructure.NewDatabase(redisConfig, os.Getenv("REDIS_PASSWORD"), os.Getenv("REDIS_DB"))
	if errRedis != nil {
		panic(errRedis.Error())
	}

	repoLoan := repositoryLoan.NewRepository(redisCache)
	serviceLoan := serviceLoan.NewService(repoLoan)

	v1 := r.Group("/api/v1")
	{
		router.Route(v1, serviceLoan)
	}

	errorHandler(r)
	r.Run(serverString)
}


//handle error method and not found endpoint
func errorHandler(r *gin.Engine) {
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, presenter.Response{
			Code:  http.StatusMethodNotAllowed,
			Message: pkg.ErrMethodNotAllow.Error(),
		})
		c.Abort()
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, presenter.Response{
			Code:    http.StatusNotFound,
			Message: pkg.ErrInvalidURL.Error(),
		})
		c.Abort()
	})
}
