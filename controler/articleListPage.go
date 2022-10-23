package controler

import (
	"blog/api"
	"blog/response"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// 查询文章列表
func ArticleListPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	articles, err := api.QueryArticleList(category)
	if err != nil {
		fmt.Printf("Query article list faild:%s", err.Error())
	}
	response.Temp.ExecuteTemplate(w, "list.html", articles)
}
