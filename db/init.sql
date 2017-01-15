CREATE TABLE snippets (
  snippet_id int(11) NOT NULL AUTO_INCREMENT,
  snippet_title varchar(255) NOT NULL,
  snippet_text text NOT NULL,
  user_id int(11) NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL,
  PRIMARY KEY (snippet_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;