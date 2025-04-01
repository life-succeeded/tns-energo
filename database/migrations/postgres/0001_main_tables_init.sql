-- +goose Up
create table if not exists roles
(
    id   int primary key generated always as identity,
    name text not null
);

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

create table if not exists objects
(
    id        int primary key generated always as identity,
    city      text not null,
    street    text not null,
    house     text not null,
    apartment text
);

create table if not exists contracts
(
    id             int primary key generated always as identity,
    number         text not null,
    contract_date  date not null,
    account_number text not null unique
);

create table if not exists consumers
(
    id                int primary key generated always as identity,
    surname           text,
    name              text,
    patronymic        text,
    legal_entity_name text,
    object_id         int references objects (id),
    contract_id       int references contracts (id)
);

create index if not exists consumers_object_id_idx on consumers (object_id);

create table if not exists devices
(
    id                   int primary key generated always as identity,
    type                 text          not null,
    number               text          not null,
    voltage              numeric(9, 2) not null,
    amperage             numeric(9, 2) not null,
    valency_before_dot   int           not null,
    valency_after_dot    int           not null,
    verification_quarter int           not null,
    verification_year    int           not null,
    accuracy_class       text          not null,
    tariffs_count        int           not null,
    deployment_place     text          not null,
    object_id            int references objects (id)
);

create index if not exists devices_object_id_idx on devices (object_id);

create table if not exists inspections
(
    id                    int primary key generated always as identity,
    resolution            text           not null,
    act_number            text           not null,
    act_date              date           not null,
    reason                text           not null,
    method                text           not null,
    seal_number           text           not null,
    action_date           timestamptz    not null,
    autumaton_seal_number text,
    device_value          numeric(11, 2) not null,
    inspector_id          int            not null references users (id),
    object_id             int            not null references objects (id),
    consumer_id           int            not null references consumers (id),
    consumer_agent_id     int            not null references consumers (id)
);

create table if not exists inspection_images
(
    inspection_id int  not null references inspections (id),
    image_id      text not null,
    primary key (inspection_id, image_id)
);

create index if not exists inspection_images_inspection_id_idx on inspection_images (inspection_id);

-- +goose Down

drop table if exists roles;
drop table if exists users;
drop table if exists objects;
drop table if exists contracts;
drop table if exists consumers;
drop table if exists devices;
drop table if exists inspections;
drop table if exists inspection_images;