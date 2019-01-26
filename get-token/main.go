package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dgrijalva/jwt-go"
	"os"
)

var jwtSecretSigningKey []byte

func main() {
	jwtSecretSigningKey = []byte(os.Getenv("JWT_SECRET_SIGNING_KEY"))

	lambda.Start(handler)
}

func handler() (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)

	return t.SignedString(jwtSecretSigningKey)
}
