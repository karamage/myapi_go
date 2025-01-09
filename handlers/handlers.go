package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"myapi/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// 1. バイトスライスを用意する
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, "fail to read Content-Length\n", http.StatusBadRequest)
		return
	}
	reqBodybuffer := make([]byte, length)

	// 2. Read メソッドでリクエストボディを読み込む
	if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
		http.Error(w, "fail to read request body\n", http.StatusBadRequest)
		return
	}
	// 3. ボディをCloseする
	defer req.Body.Close()

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	resString := fmt.Sprintf("Article List (page %d)\n", page)
	io.WriteString(w, resString)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No. %d\n", articleID)
	io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
