package main

import (
	"blog/handlers"
	"blog/mysql"
	"html/template"
	"log"
	"net/http"
	"time"

	"blog/response"

	"github.com/gorilla/mux"
)

func init() {
	response.Temp = template.Must(template.ParseGlob("./views/*.html"))
	mysql.Connect(&mysql.Information{
		UserName:     "yangjiafeng",
		Password:     ".%fuoa;k$@*..,!Ujf",
		Addr:         "192.168.2.194:13306",
		DatabaseName: "blog",
	})
}

// 页面重定向
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/blog", http.StatusFound)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RedirectHandler)
	r.HandleFunc("/blog", handlers.SelectArticleList)
	r.HandleFunc("/blog/{id:[0-9]+}", handlers.SelectArticleContent)
	r.HandleFunc("/blog/{category}", handlers.SelectArticleList)
	r.HandleFunc("/blog/{category}", handlers.SelectArticleList)
	r.HandleFunc("/blog/{category}", handlers.SelectArticleList)

	//接口
	r.HandleFunc("/blog/api/insert", handlers.InsertArticle)
	fileServer := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/").Handler(http.StripPrefix("/static", fileServer))
	s := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
