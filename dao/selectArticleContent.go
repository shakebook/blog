package dao

import (
	"blog/mysql"
	"blog/response"
)

// 查询文章内容
func SelectArticleContent(id int) (*response.ArticleContent, error) {
	db := mysql.GetDB()
	sql := `SELECT 
	id,
	title,
	content
	FROM article_content WHERE id = ? limit 1`

	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var content *response.ArticleContent
	err = stmt.QueryRow(id).Scan(
		&content.Id,
		&content.Title,
		&content.Content,
	)
	if err != nil {
		return nil, err
	}

	return content, nil
}
