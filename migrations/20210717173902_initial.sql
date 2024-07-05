-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
   id BIGSERIAL NOT NULL,

   CONSTRAINT pk_user_id PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
