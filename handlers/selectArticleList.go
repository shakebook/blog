package handlers

import (
	"blog/dao"
	"blog/response"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// 查询文章列表
func SelectArticleList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	articles, err := dao.SelectArticleList(category)
	fmt.Printf("select articles, %v", articles)
	if err != nil {
		fmt.Printf("Query article list faild:%s", err.Error())
	}
	response.Temp.ExecuteTemplate(w, "list.html", articles)
}
