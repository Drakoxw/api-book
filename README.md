# API BOOKS - GO

## COMMANDS

#### Go Run
```sh
    go run main.go
```
#### Build Image Docker
```sh
    docker build -t api-book:latest .
```

#### Build Image Docker
```sh
    docker run -d -p 3300:3000 --name api-book-app api-book
```

## AWS make lambda
```sh
    GOOS=linux GOARCH=amd64 go build -o main
```