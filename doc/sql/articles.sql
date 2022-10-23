CREATE TABLE articles
(
   id bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
   title varchar(100) NOT NULL COMMENT '标题',
   author varchar(20) NOT NULL COMMENT '作者',
   category varchar(20) NOT NULL COMMENT '类型',
   category_name varchar(20) NOT NULL COMMENT '类型名称',
   image_url varchar(100) NOT NULL COMMENT '封面',
   browser_number bigint(20) DEFAULT 0 COMMENT '浏览量',
   thumbs_number bigint(20) DEFAULT 0 COMMENT '点赞数',
   create_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
   update_time datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
   PRIMARY KEY(id),
   UNIQUE KEY index_title (title) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

ALTER TABLE articles ADD COLUMN content text AFTER image_url;
