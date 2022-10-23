package controler

import (
	"blog/api"
	"blog/response"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// 新增文章
func EditArticleContent(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		res(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var article *response.Article
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&article)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			res(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			res(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	err = api.EditArticleContent(article)
	if err != nil {
		fmt.Printf("Insert article failed, %s", err.Error())
		res(w, "Insert article failed, "+err.Error(), http.StatusInternalServerError)
		return
	}
	res(w, "Success", http.StatusOK)
}
