build:
	cd example-api && GOOS=linux go build
	cd verify-token && GOOS=linux go build
	cd get-token && GOOS=linux go build

local: build
	sam local start-api

deploy: check-jwt-secret-signing-key build
	aws cloudformation package \
		--template-file template.yml \
		--output-template-file packaged-template.yml \
		--s3-bucket markwilson-authorisedlambdaapi
	aws cloudformation deploy \
		--template-file packaged-template.yml \
		--stack-name authorised-lambda-api \
		--parameter-overrides JWTSecretSigningKey="${JWT_SECRET_SIGNING_KEY}" \
		--capabilities CAPABILITY_IAM

destroy-service:
	aws cloudformation delete-stack --stack-name authorised-lambda-api

.PHONY: get-token
get-token:
	@${MAKE} get-lambda-name-for-get-token | xargs -I % aws lambda invoke --function-name % out.txt > /dev/null
	@echo "Token: `cat out.txt`"
	@rm out.txt

call-api: check-token
	curl -H "Authorization: ${TOKEN}" -i `${MAKE} get-deployed-url`

get-deployed-url:
	@aws cloudformation describe-stacks --stack-name authorised-lambda-api | jq '.Stacks[0] .Outputs[] | select(.OutputKey == "SensitiveFunctionUrl") | .OutputValue' -r

get-lambda-name-for-get-token:
	@aws cloudformation describe-stacks --stack-name authorised-lambda-api | jq '.Stacks[0] .Outputs[] | select(.OutputKey == "GetTokenLambda") | .OutputValue' -r

check-jwt-secret-signing-key:
ifndef JWT_SECRET_SIGNING_KEY
	$(error "JWT_SECRET_SIGNING_KEY is missing")
endif

check-token:
ifndef TOKEN
	$(error "TOKEN is missing")
endif