CREATE TABLE article_content
(
   id bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
   article_id bigint(20) NOT NULL COMMENT '文章ID',
   content text NOT NULL COMMENT '内容',
   PRIMARY KEY(id),
   UNIQUE KEY article_id (article_id) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
ALTER TABLE article_content ADD PRIMARY KEY (id);
ALTER TABLE article_content DROP PRIMARY KEY
ALTER TABLE article_content ADD UNIQUE (article_id);

ALTER TABLE article_content DROP CONSTRAINT article_id;

ALTER TABLE article_content modify id bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键';
