CREATE TABLE article_content
(
   id bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
   article_id bigint(20) NOT NULL COMMENT '文章ID',
   title varchar(100) NOT NULL COMMENT '标题',
   content text NOT NULL COMMENT '内容',
   PRIMARY KEY(id, article_id),
   UNIQUE KEY index_title (title) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
