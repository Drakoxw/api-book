package aws

import (
	"api-book/internal/domain/repository"
	"api-book/internal/infrastructure/pkg"
	"api-book/internal/infrastructure/utils"
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func HandleListUser(ctx context.Context, ev events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res := utils.CreateAwsResponse()
	db := pkg.StartDB()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	users, _ := userRepo.FindAllUsers()

	resData, err := utils.CreateResponseApi(users)
	if err != nil {
		res.Body = err.Error()
		res.StatusCode = http.StatusUnprocessableEntity
		return res, nil
	}

	res.Body = string(resData)
	res.StatusCode = http.StatusOK

	return res, nil
}
