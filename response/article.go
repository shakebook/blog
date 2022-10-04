package response

import (
	"html/template"
)

var Temp *template.Template

type Article struct {
	Id            int
	Title         string
	Author        string
	Category      string
	CategoryName  string
	ImageUrl      string
	BrowserNumber int
	ThumbsNumber  int
	CreateTime    string
	UpdateTime    string
	Content       string
}

type ArticleContent struct {
	Id      int
	Title   string
	Content string
}
