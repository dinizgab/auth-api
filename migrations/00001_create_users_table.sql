-- +goose Up
-- +goose StatementBegin
create table if not exists users (
    id uuid default gen_random_uuid(),
    username varchar(100),
    email varchar(255),
    password varchar(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
