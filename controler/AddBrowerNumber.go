package controler

import (
	"blog/api"
	"fmt"
	"net/http"
	"strconv"
)

func AddBrowerNumber(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id, _ := strconv.Atoi(values.Get("articleId"))
	err := api.AddBrowerNumber(id)
	if err != nil {
		fmt.Printf("Create article brower number faild:%s", err.Error())
	}
	res(w, "success", http.StatusOK)
}
