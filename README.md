# Sample application for Go Web Services

A very basic web-services application built with go-lang and the Gin Web Framework.

Based on https://github.com/MBvisti/weight-tracker-article.

## Database

Uses a locally running database, which you won't have, so that might be a bit tricky for you.

## Dependencies

```
go get go.mongodb.org/mongo-driver/mongo
go get -u github.com/gin-gonic/gin
go get github.com/rs/cors/wrapper/gin
```

## Running

```
go run ./cmd/server/main.go
```

## Building

```
go build ./cmd/server/main.go                       will create a file named "main"
go build -o thing-server ./cmd/server/main.go       will create a file named "thing-server"
```
