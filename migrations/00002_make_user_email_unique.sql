-- +goose Up
-- +goose StatementBegin
alter table users add constraint email unique (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table users drop constraint email;
-- +goose StatementEnd
