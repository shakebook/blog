package controler

import (
	"blog/api"
	"blog/response"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// 查询文章
func ArticleContentPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	articleId, _ := strconv.Atoi(id)
	articleContent, err := api.QueryArticleContent(articleId)
	if err != nil {
		fmt.Printf("Query article content faild:%s", err.Error())
	}
	response.Temp.ExecuteTemplate(w, "article.html", articleContent)
}
