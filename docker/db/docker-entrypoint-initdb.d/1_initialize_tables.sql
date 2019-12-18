CREATE TABLE customers
(
	id int unsigned NOT NULL AUTO_INCREMENT,
	created_at datetime NOT NULL COMMENT '作成日時',
	modified_at datetime NOT NULL COMMENT '更新日時',
	PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT = '会員';