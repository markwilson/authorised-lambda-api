# Authorised Lambda API

This is an example of a self-contained SAM configuration for handling custom authorisation of API Gateways.

## Setup

1. Clone this repository
2. `export JWT_SECRET_SIGNING_KEY="... some secret key..."`
2. `export S3_BUCKET="... some manually created bucket for uploading your code into..."`
3. `make deploy`

## Using the deployed service

### Getting a token

```sh
make get-token
# Outputs "Token: ..."
```

### Call the API

```sh
TOKEN="..." make call-api
```

## Cleanup

`make destroy-service` deletes the deployed stack.