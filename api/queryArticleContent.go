package api

import (
	"blog/mysql"
	"blog/response"
	"database/sql"
	"encoding/json"
	"html/template"
)

// 查询文章内容
func QueryArticleContent(articleId int) (*response.Article, error) {
	db := mysql.GetDB()
	query := `SELECT 
	article_content.content,
	articles.title,
	articles.image_url
	FROM article_content 
	LEFT JOIN articles ON article_content.article_id=articles.id 
	WHERE article_content.article_id = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// rows
	var rows *sql.Rows
	rows, err = stmt.Query(articleId)
	if err != nil {
		return nil, err
	}

	res := []*response.Article{}
	for rows.Next() {
		item := &response.Article{}
		var contentString string
		var content []response.Content
		rows.Scan(
			&contentString,
			&item.Title,
			&item.ImageUrl,
		)
		if err = json.Unmarshal([]byte(contentString), &content); err != nil {
			return nil, err
		}
		for k, _ := range content {
			content[k].Paragraph = template.HTML(content[k].Paragraph)
		}
		item.Content = content

		item.Id = articleId
		res = append(res, item)
	}

	if err != nil {
		return nil, err
	}

	return res[0], nil
}
