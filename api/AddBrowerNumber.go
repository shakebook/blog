package api

import (
	"blog/mysql"
)

func AddBrowerNumber(articleId int) error {
	db := mysql.GetDB()
	sql := `UPDATE articles SET browser_number=browser_number+1 WHERE id=?`
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(articleId)
	if err != nil {
		return err
	}
	return nil
}
