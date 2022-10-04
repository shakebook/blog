package dao

import (
	"blog/mysql"
	"blog/response"
	"blog/tools"
	"database/sql"
	"time"
)

// 查询文章列表
func SelectArticleList(category string) ([]*response.Article, error) {
	db := mysql.GetDB()
	query := `SELECT 
	id,
	title,
	author,
	category,
	category_name,
	image_url,
	browser_number,
	thumbs_number,
	create_time,
	update_time
	FROM articles`
	if category != "" {
		query = tools.PingString([]string{query, ` WHERE category = ?`})
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// rows
	var rows *sql.Rows
	if category != "" {
		rows, err = stmt.Query(category)
		if err != nil {
			return nil, err
		}
	} else {
		rows, err = stmt.Query()
		if err != nil {
			return nil, err
		}
	}

	res := []*response.Article{}
	for rows.Next() {
		item := &response.Article{}
		var createTime *time.Time
		var updateTime *time.Time
		rows.Scan(
			&item.Id,
			&item.Title,
			&item.Author,
			&item.Category,
			&item.CategoryName,
			&item.ImageUrl,
			&item.BrowserNumber,
			&item.ThumbsNumber,
			&createTime,
			&updateTime,
		)
		item.CreateTime = createTime.Format("2006-01-02 15:04:05")
		item.UpdateTime = updateTime.Format("2006-01-02 15:04:05")
		res = append(res, item)
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}
