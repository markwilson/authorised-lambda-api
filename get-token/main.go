// package main generates valid JWTs for authorising the API
package main

import (
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

// handler generates a JWT and sends it back to the invoker
func handler() (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)

	return t.SignedString(jwtSecretSigningKey)
}
