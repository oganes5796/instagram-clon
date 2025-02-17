-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id serial primary key,
    email varchar(255) not null,
    username varchar(255) not null,
    password varchar(255) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
