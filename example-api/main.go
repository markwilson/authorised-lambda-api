// package main contains a basic API Gateway->Lambda handler with assumed user authorisation
package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

// main starts the Lambda handler
func main() {
	lambda.Start(handler)
}

// handler receives API Gateway requests and returns a response dynamic to the user
func handler(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	output, ok := r.RequestContext.Authorizer["principalId"].(string)
	if !ok {
		output = "unknown"
	}

	return events.APIGatewayProxyResponse{
		Body:       "My sensitive content: " + output,
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
	}, nil
}
