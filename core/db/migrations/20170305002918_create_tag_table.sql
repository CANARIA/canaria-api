
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tags (
  tag_id bigint NOT NULL AUTO_INCREMENT,
  tag_name varchar(255) NOT NULL UNIQUE,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL DEFAULT now(),
  PRIMARY KEY (tag_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE tags
