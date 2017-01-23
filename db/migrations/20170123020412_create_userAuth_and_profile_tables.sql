
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE account (
  user_id bigint NOT NULL AUTO_INCREMENT,
  user_name varchar(255) NOT NULL,
  mailaddress varchar(255) NOT NULL,
  password int(11) NOT NULL,
  roll tinyint(1) NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL,
  is_deleted tinyint(1) NOT NULL,
  PRIMARY KEY (user_id),
  UNIQUE (user_name, mailaddress)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE profile (
  display_name varchar(255) NOT NULL,
  bio text,
  url varchar(255),
  user_id bigint NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL,
  is_deleted tinyint(1) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES account(user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE profile;
DROP TABLE account;
