Go Fetcher Server and Client
==========================

**Stack**:

- GO 1.20.4
- MongoDB 6.0.5
- gRPC

This is a simple fetcher service which fetch price data from URL and store it in MongoDB

Please check [Proto](proto) for API details

## Before run:

1. Setup you configs for MongoDB. You can check example.env

## To run server:

```shell
go run cmd/server/main.go
```

## To run client example:

```shell
go run cmd/client/main.go
```
