package main

import (
	/*
		"log"
		"myapi/handlers"
		"net/http"
		"github.com/gorilla/mux"
	*/
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "root"
	dbPassword := "pass"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@/%s", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Successfully connected to the database")
	}

	/*
		r := mux.NewRouter()

		r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
		r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
		r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
		r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
		r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
		r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

		log.Print("server start at port 8080")
		log.Fatal(http.ListenAndServe(":8080", r))
	*/
}
