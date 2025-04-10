-- +goose Up
create table if not exists roles
(
    id   int primary key generated always as identity,
    name text not null
);

insert into roles (name)
values ('Инспектор'),
       ('Админ');

create table if not exists users
(
    id                       int primary key generated always as identity,
    email                    text        not null unique,
    surname                  text        not null,
    name                     text        not null,
    patronymic               text        not null,
    position                 text        not null,
    password_hash            text        not null,
    refresh_token            text        not null,
    refresh_token_expires_at timestamptz not null,
    role_id                  int         not null references roles (id)
);

-- +goose Down

drop table if exists roles;
drop table if exists users;