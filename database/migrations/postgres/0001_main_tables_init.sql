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
    email                    text not null unique,
    surname                  text not null,
    name                     text not null,
    patronymic               text,
    position                 text,
    password_hash            text,
    refresh_token            text,
    refresh_token_expires_at timestamptz,
    role_id                  int  not null references roles (id)
);

-- +goose Down

drop table if exists roles;
drop table if exists users;