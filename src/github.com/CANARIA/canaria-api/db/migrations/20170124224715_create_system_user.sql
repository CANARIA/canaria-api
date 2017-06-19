
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
START TRANSACTION;

insert into
  accounts(user_name, mailaddress, password)
  values("system", "canaria", "password");

insert into
  profiles(display_name, user_id)
  values("system", 1);

COMMIT;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

