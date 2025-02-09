package main

import (
	"log"
	"myapi/controllers"
	"myapi/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var service *services.MyAppService // TODO
	con := controllers.NewMyAppController(service)

	r := mux.NewRouter()

	r.HandleFunc("/hello", con.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	log.Print("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
