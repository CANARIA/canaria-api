
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
START TRANSACTION;

CREATE TABLE accounts (
  user_id bigint NOT NULL AUTO_INCREMENT,
  user_name varchar(255) NOT NULL,
  mailaddress varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  roll tinyint(1) NOT NULL DEFAULT 1,
  created_at datetime NOT NULL DEFAULT now(),
  updated_at datetime NOT NULL DEFAULT now(),
  is_deleted tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (user_id),
  UNIQUE (user_name, mailaddress)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE profiles (
  display_name varchar(255) NOT NULL,
  bio text,
  url varchar(255),
  user_id bigint NOT NULL,
  created_at datetime NOT NULL DEFAULT now(),
  updated_at datetime NOT NULL DEFAULT now(),
  is_deleted tinyint(1) NOT NULL DEFAULT 0,
  FOREIGN KEY (user_id) REFERENCES accounts(user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

COMMIT;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE profile;
DROP TABLE account;
