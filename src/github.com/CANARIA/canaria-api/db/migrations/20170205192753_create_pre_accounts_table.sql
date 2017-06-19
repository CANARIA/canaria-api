
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE pre_accounts (
  id bigint NOT NULL AUTO_INCREMENT,
  url_token varchar(255) NOT NULL,
  created_at datetime NOT NULL,
  mailaddress varchar(255) NOT NULL,
  is_registered tinyint(1) DEFAULT 0,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back


DROP TABLE pre_accounts;
