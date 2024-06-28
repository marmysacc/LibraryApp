package main

import (
	"fmt"
	_ "library-app/docs"
	"library-app/handler"
	"library-app/model"
	"library-app/repository"
	"library-app/service"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Library API
// @version 1.0
// @description This is a sample server for managing a library.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	dsn := "user=CarRental password=CarRental dbname=LibraryApp port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&model.Book{}, &model.Author{})

	bookRepo := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	authorRepo := repository.NewAuthorRepository(db)
	authorService := service.NewAuthorService(authorRepo)
	authorHandler := handler.NewAuthorHandler(authorService)

	r := mux.NewRouter()

	r.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", bookHandler.GetBookByID).Methods("GET")
	r.HandleFunc("/books", bookHandler.GetAllBooks).Methods("GET")
	r.HandleFunc("/books/{id}", bookHandler.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", bookHandler.DeleteBook).Methods("DELETE")

	r.HandleFunc("/authors", authorHandler.CreateAuthor).Methods("POST")
	r.HandleFunc("/authors/{id}", authorHandler.GetAuthorByID).Methods("GET")
	r.HandleFunc("/authors", authorHandler.GetAllAuthors).Methods("GET")
	r.HandleFunc("/authors/{id}", authorHandler.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id}", authorHandler.DeleteAuthor).Methods("DELETE")

	// Swagger endpoint
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
