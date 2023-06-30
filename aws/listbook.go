package aws

import (
	"api-book/internal/domain/repository"
	"api-book/internal/infrastructure/pkg"
	"api-book/internal/infrastructure/utils"
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func HandleListBooks(ctx context.Context, ev events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res := utils.CreateAwsResponse()
	db := pkg.StartDB()
	defer db.Close()

	page, limit := utils.GetPaginatorAwsSql(ev.Headers)

	bookRepo := repository.NewBookRepository(db)
	books, _ := bookRepo.ListBooks(page, limit)

	resData, err := utils.CreateResponseApi(books)
	if err != nil {
		res.Body = err.Error()
		res.StatusCode = http.StatusUnprocessableEntity
		return res, nil
	}

	res.Body = string(resData)
	res.StatusCode = http.StatusOK

	return res, nil
}
