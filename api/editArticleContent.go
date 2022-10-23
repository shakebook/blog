package api

import (
	"blog/mysql"
	"blog/response"
	"encoding/json"
	"errors"
)

func EditArticleContent(r *response.Article) error {
	db := mysql.GetDB()
	sql := `INSERT INTO article_content (article_id,content) VALUES (?,?) ON DUPLICATE KEY UPDATE content=?`
	//sql := `REPLACE INTO article_content (article_id, content) VALUES (?,?)`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	content, _ := json.Marshal(r.Content)
	res, err := stmt.Exec(r.Id, string(content), string(content))
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
