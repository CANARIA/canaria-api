
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE popular_tags (
  tag_id bigint NOT NULL,
  tag_name varchar(255) NOT NULL,
  created_at datetime NOT NULL DEFAULT now(),
  PRIMARY KEY (tag_id),
  FOREIGN KEY (tag_id) REFERENCES tags(tag_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE popular_tags
