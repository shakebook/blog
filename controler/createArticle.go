package controler

import (
	"blog/api"
	"blog/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/google/uuid"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("imageUrl")
	title := r.FormValue("title")
	author := r.FormValue("author")
	category := r.FormValue("category")
	categoryName := r.FormValue("categoryName")
	if err != nil {
		res(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}
	defer file.Close()
	suffix := path.Ext(handler.Filename)
	imageName := uuid.New().String() + suffix

	f, err := os.OpenFile("./web/images/"+imageName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		res(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}
	defer f.Close()
	_, _ = io.WriteString(w, "File: "+handler.Filename+" Uploaded successfully\n")
	_, _ = io.Copy(f, file)

	err = api.CreateArticle(&response.Article{
		Title:        title,
		Author:       author,
		Category:     category,
		CategoryName: categoryName,
		ImageUrl:     imageName,
	})
	if err != nil {
		fmt.Printf("Insert article failed, %s", err.Error())
		res(w, "Insert article failed, "+err.Error(), http.StatusInternalServerError)
		return
	}
	res(w, "Success", http.StatusOK)
}

func res(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
