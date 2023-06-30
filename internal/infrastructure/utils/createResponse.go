package utils

import (
	"api-book/internal/domain/dtos"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func CreateAwsResponse() events.APIGatewayProxyResponse {
	res := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "GET,HEAD,OPTIONS,POST",
		},
	}

	return res
}

func CreateResponseApi(data interface{}) ([]byte, error) {
	responseDto := dtos.ResponseDTO{
		Status:  "success",
		Message: "data found",
		Data:    data,
	}

	return json.Marshal(responseDto)
}
