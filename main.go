package main

import (
	"api-book/aws"
	"api-book/internal/domain/repository"
	api "api-book/internal/infrastructure/api/user"
	"api-book/internal/infrastructure/pkg"
	"api-book/internal/infrastructure/utils"
	"fmt"

	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

func startServer() {
	db := pkg.StartDB()
	defer db.Close()
	port := utils.GetPort()
	userRepo := repository.NewUserRepository(db)
	userHandler := &api.UserHandler{UserRepo: userRepo}

	http.HandleFunc("/users", userHandler.GetAllUsers)
	fmt.Println("Run on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func CompileListUser() {
	lambda.Start(aws.HandleListUser)
}

func main() {
	// compileListUser()
	startServer()
}
