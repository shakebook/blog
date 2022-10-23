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
	Content       []Content
}

type Content struct {
	Paragraph template.HTML
	PreCode   string
	Image     string
}
