package handlers

import (
	"fmt"
	"io"
	"net/http"
)

//  Define handler
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// Handler processing details
	http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w , "Posting Article...\n")
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w , "Article List...\n")
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID :=1
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}