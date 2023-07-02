package main

import (
	"api-book/aws"
	"api-book/internal/domain/repository"
	api "api-book/internal/infrastructure/api"
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
	bookRepo := repository.NewBookRepository(db)
	lendRepo := repository.NewLendBookRepository(db)

	userHandler := &api.UserHandler{UserRepo: userRepo}
	bookHandler := &api.BookHandler{BookRepo: bookRepo}
	lendHandler := &api.LendBookHandler{LendBookRepo: lendRepo}

	http.HandleFunc("/user", userHandler.GetUserId)
	http.HandleFunc("/users", userHandler.GetAllUsers)
	http.HandleFunc("/new_user", userHandler.CreateUser)
	http.HandleFunc("/update_user", userHandler.UpdateUser)

	http.HandleFunc("/book", bookHandler.GetBook)
	http.HandleFunc("/books", bookHandler.ListBooks)
	http.HandleFunc("/new_book", bookHandler.CreateBook)

	http.HandleFunc("/new_book_loan", lendHandler.CreateLendBook)
	http.HandleFunc("/return_book", lendHandler.ReturnBookToLibrary)

	http.HandleFunc("/history_books", bookHandler.GetHistoryLendBook)
	http.HandleFunc("/history_users", userHandler.GetHistoryLendUsers)

	utils.LogInfo("main.log", fmt.Sprintf("Server start %s", port))
	log.Fatal(http.ListenAndServe(port, nil))
}

func compileListUser() {
	lambda.Start(aws.HandleListUser)
}
func compileListBooks() {
	lambda.Start(aws.HandleListBooks)
}

func main() {
	// compileListBooks()
	// compileListUser()
	startServer()
}
