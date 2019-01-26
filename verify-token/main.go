// package main verifies the provided JWT is valid
package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dgrijalva/jwt-go"
	"os"
)

// jwtSecretSigningKey is used for by jwt-go and must be kept secret
var jwtSecretSigningKey []byte

// main starts the Lambda handler
func main() {
	jwtSecretSigningKey = []byte(os.Getenv("JWT_SECRET_SIGNING_KEY"))

	lambda.Start(handler)
}

// handler receives the API Gateway custom authorizer request and checks it is valid
func handler(r events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	token, err := jwt.Parse(r.AuthorizationToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecretSigningKey, nil
	})

	if err == nil && token.Valid {
		return events.APIGatewayCustomAuthorizerResponse{
			PrincipalID: "test user",
			PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
				Version: "2012-10-17",
				Statement: []events.IAMPolicyStatement{{
					Action:   []string{"execute-api:Invoke"},
					Effect:   "Allow",
					Resource: []string{r.MethodArn},
				}},
			},
		}, nil
	}

	return events.APIGatewayCustomAuthorizerResponse{
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{{
				Action:   []string{"execute-api:Invoke"},
				Effect:   "Deny",
				Resource: []string{r.MethodArn},
			}},
		},
	}, nil
}
