package main

import (
	"blog/auth"
	"blog/controler"
	"blog/mysql"
	"blog/response"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// 初始化
func init() {
	response.Temp = template.Must(template.ParseGlob("./web/views/*.html"))
	mysql.Connect(&mysql.Infor{
		UserName:     "root",
		Password:     ".",
		Addr:         "",
		DatabaseName: "blog",
	})
}

// 页面重定向
func HomePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/blog", http.StatusFound)
}

// 拦截器
func interceptor(next http.Handler) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		valid := auth.Done(w, r)
		if valid {
			next.ServeHTTP(w, r)
		}
	})
}

// 主程
func main() {
	r := mux.NewRouter()

	//路由器
	r.HandleFunc("/", HomePage)
	r.HandleFunc("/blog", controler.ArticleListPage)
	r.HandleFunc("/blog/{id:[0-9]+}", controler.ArticleContentPage)
	r.HandleFunc("/blog/{category}", controler.ArticleListPage)

	//接口
	//文章
	r.HandleFunc("/blog/api/article", interceptor(http.HandlerFunc(controler.CreateArticle))).Methods("POST")
	r.HandleFunc("/blog/api/article/content", interceptor(http.HandlerFunc(controler.EditArticleContent))).Methods("POST")
	r.HandleFunc("/blog/api/article/brower", interceptor(http.HandlerFunc(controler.AddBrowerNumber))).Methods("GET")

	//文件服务
	fileServer := http.FileServer(http.Dir("./web/"))
	r.PathPrefix("/").Handler(http.StripPrefix("/web", fileServer))

	s := &http.Server{
		Handler:      r,
		Addr:         ":80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
