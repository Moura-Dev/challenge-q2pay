# Challenge Q2Pay

## Architecture

![img](docs/img/dbdiagram.png)


## Docs
Docs: http://localhost:5000/api/docs/index.html#/

# Golang Api

## Requirements
- Docker
- Golang

## Download Collection Postman
docs/quero2pay.postman_collection.json

## Start postgresdb
```bash
make services
```

## Run

Go mod Init

```bash
go mod init challenge-q2pay
```


Install Dependencies

```bash
  go mod tidy
```

Run Server Service

```bash
 make run
```
