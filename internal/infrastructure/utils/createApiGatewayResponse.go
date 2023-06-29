package utils

import "github.com/aws/aws-lambda-go/events"

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
