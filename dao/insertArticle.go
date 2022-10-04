package dao

import (
	"blog/mysql"
	"blog/response"
	"errors"
)

// 查询文章列表
func InsertArticle(r *response.Article) error {
	db := mysql.GetDB()
	sql := `INSERT IGNORE INTO articles (title, author, category, category_name, image_url, content) VALUES (?,?,?,?,?,?)`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(r.Title, r.Author, r.Category, r.CategoryName, r.ImageUrl, r.Content)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("existing")
	}

	return nil
}
