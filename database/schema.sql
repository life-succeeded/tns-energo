create table if not exists users
(
    id                          int primary key generated always as identity,
    role_id                     int         not null, -- 0 - Inspector, 1 - Dispatcher, 2 - Specialist
    surname                     text        not null,
    name                        text        not null,
    patronymic                  text,
    phone_number                text        not null unique check ( phone_number ~ '^(\+7|8)\d{10}$' ),
    email                       text        not null unique check ( email ~ '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$' ),
    password_hash               text,
    refresh_token               text,
    refresh_token_expired_after timestamptz,
    created_at                  timestamptz not null default now(),
    updated_at                  timestamptz not null default now()
);

create table if not exists brigade_statuses
(
    id   int primary key generated always as identity,
    name text not null
);

insert into brigade_statuses (name)
values ('Idle'),
       ('OnTask'),
       ('Archived');

create table if not exists brigades
(
    id         int primary key generated always as identity,
    status     int         not null references brigade_statuses (id) on delete restrict,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

create table if not exists brigade_members
(
    brigade_id   int         not null references brigades (id) on delete cascade,
    inspector_id int         not null references users (id) on delete cascade,
    assigned_at  timestamptz not null default now(),
    primary key (brigade_id, inspector_id)
);

create table if not exists subscriber_statuses
(
    id   int primary key generated always as identity,
    name text not null
);

insert into subscriber_statuses (name)
values ('Active'),
       ('Violator'),
       ('Archived');

create table if not exists subscribers
(
    id             int primary key generated always as identity,
    account_number text        not null unique,                        -- Лицевой счет
    surname        text        not null,
    name           text        not null,
    patronymic     text        not null,
    phone_number   text        not null check ( phone_number ~ '^(\+7|8)\d{10}$' ),
    email          text        not null check ( email ~ '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$' ),
    inn            text        not null check ( inn ~ '^\d{10,12}$' ), -- ИНН
    birth_date     date        not null check ( '1900-01-01' <= birth_date and birth_date < current_date ),
    status         int         not null references subscriber_statuses (id) on delete restrict,
    created_at     timestamptz not null default now(),
    updated_at     timestamptz not null default now()
);

create table if not exists passports
(
    id            int primary key generated always as identity,
    subscriber_id int  not null references subscribers (id) on delete cascade,
    series        text not null check ( series ~ '^\d{4}$' ), -- Серия
    number        text not null check ( number ~ '^\d{6}$' ), -- Номер
    issued_by     text not null,                              -- Кем выдан
    issue_date    date not null                               -- Когда выдан
);

create table if not exists objects
(
    id             int primary key generated always as identity,
    address        text        not null,
    have_automaton bool        not null, -- Наличие коммутационного (вводного) аппарата
    created_at     timestamptz not null default now(),
    updated_at     timestamptz not null default now()
);

create table if not exists device_place_types
(
    id   int primary key generated always as identity,
    name text not null
);

insert into device_place_types (name)
values ('Other'),
       ('Flat'),
       ('StairLanding');

create table if not exists devices
(
    id                int primary key generated always as identity,
    object_id         int         not null references objects (id) on delete cascade,
    type              text        not null,
    number            text        not null,
    place_type        int         not null references device_place_types (id) on delete restrict,
    place_description text        not null, -- Место установки прибора учета
    created_at        timestamptz not null default now(),
    updated_at        timestamptz not null default now()
);

create table if not exists seals
(
    id         int primary key generated always as identity,
    device_id  int         not null references devices (id) on delete cascade,
    number     text        not null,
    place      text        not null, -- Место установки
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

create table if not exists contracts
(
    id            int primary key generated always as identity,
    number        text        not null unique,
    subscriber_id int         not null references subscribers (id) on delete restrict,
    object_id     int         not null references objects (id) on delete restrict,
    sign_date     date        not null check ( sign_date <= current_date ), -- Дата подписания
    created_at    timestamptz not null default now(),
    updated_at    timestamptz not null default now()
);

create table if not exists task_statuses
(
    id   int primary key generated always as identity,
    name text not null
);

insert into task_statuses (name)
values ('Planned'),
       ('InWork'),
       ('Done');

create table if not exists tasks
(
    id            int primary key generated always as identity,
    brigade_id    int         references brigades (id) on delete set null,
    object_id     int         not null references objects (id) on delete cascade,
    plan_visit_at timestamptz, -- Запланированное время визита
    status        int         not null references task_statuses (id) on delete restrict,
    comment       text,
    started_at    timestamptz,
    finished_at   timestamptz,
    created_at    timestamptz not null default now(),
    updated_at    timestamptz not null default now(),
    check ( started_at is null or finished_at is null or started_at <= finished_at )
);

create table if not exists storage_files
(
    id        int primary key generated always as identity,
    file_name text not null,
    file_size int  not null,
    bucket    text not null,
    url       text not null -- Ссылка на скачивание
);

create table if not exists inspection_statuses
(
    id   int primary key generated always as identity,
    name text not null
);

insert into inspection_statuses (name)
values ('InWork'),
       ('Done');

create table if not exists inspection_types
(
    id   int primary key generated always as identity,
    name text not null
);

insert into inspection_types (name)
values ('Limitation'),
       ('Resumption'),
       ('Verification'),
       ('UnauthorizedConnection');

create table if not exists inspection_resolutions
(
    id   int primary key generated always as identity,
    name text not null
);

insert into inspection_resolutions (name)
values ('Limited'),
       ('Stopped'),
       ('Resumed');

create table if not exists inspection_methods_by
(
    id   int primary key generated always as identity,
    name text not null
);

insert into inspection_methods_by (name)
values ('Consumer'),
       ('Inspector');

create table if not exists inspection_reason_types
(
    id   int primary key generated always as identity,
    name text not null
);

insert into inspection_reason_types (name)
values ('NotIntroduced'),
       ('ConsumerLimited'),
       ('InspectorLimited'),
       ('Resumed');

create table if not exists inspections
(
    id                        int primary key generated always as identity,
    task_id                   int         not null unique references tasks (id) on delete restrict,
    status                    int         not null references inspection_statuses (id) on delete restrict,
    type                      int references inspection_types (id) on delete restrict,
    resolution                int references inspection_resolutions (id) on delete restrict, -- Результаты проверки
    limit_reason              text,                                                          -- Основание введения ограничения (приостановления) режима потребления. Если NULL, то это неполная оплата
    method                    text,                                                          -- Способ введения ограничения, приостановления, возобновления режима потребления, номера и место установки пломб
    method_by                 int references inspection_methods_by (id) on delete restrict,  -- Кем было введено
    reason_type               int references inspection_reason_types (id) on delete restrict,
    reason_description        text,                                                          -- Причина невведения ограничения (только для reason_type = 1)
    is_restriction_checked    bool,                                                          -- Произведена проверка введенного ограничения
    is_violation_detected     bool,                                                          -- Нарушение потребителем введенного ограничения выявлено
    is_expense_available      bool,                                                          -- Наличие расхода после введенного ограничения
    violation_description     text,                                                          -- Иное описание выявленного нарушения/сведения, на основании которых сделан вывод о нарушении
    is_unauthorized_consumers bool,                                                          -- Самовольное подключение энергопринимающих устройств Потребителя к электрическим сетям
    unauthorized_description  text,                                                          -- Описание места и способа самовольного подключения к электрическим сетям
    unauthorized_explanation  text,                                                          -- Объяснение лица, допустившего самовольное подключение к электрическим сетям
    inspect_at                timestamptz,                                                   -- Дата проверки
    energy_action_at          timestamptz,                                                   -- Время действия над подачей электроэнергии
    created_at                timestamptz not null default now(),
    updated_at                timestamptz not null default now()
);

create table if not exists inspected_devices
(
    id            int primary key generated always as identity,
    device_id     int            not null references devices (id) on delete cascade,
    inspection_id int            not null references inspections (id) on delete cascade,
    value         numeric(15, 2) not null, -- Текущее показание
    consumption   numeric(15, 2) not null, -- Расход электрической энергии кВтч
    created_at    timestamptz    not null default now()
);

create table if not exists inspected_seals
(
    id            int primary key generated always as identity,
    seal_id       int         not null references seals (id) on delete cascade,
    inspection_id int         not null references inspections (id) on delete cascade,
    is_broken     bool        not null, -- Сорвана ли пломба
    created_at    timestamptz not null default now()
);

create table if not exists attachment_types
(
    id   int primary key generated always as identity,
    name text not null
);

insert into attachment_types (name)
values ('DevicePhoto'),
       ('SealPhoto'),
       ('Act');

create table if not exists inspection_attachments
(
    id            int primary key generated always as identity,
    inspection_id int         not null references inspections (id) on delete cascade,
    type          int         not null references attachment_types (id) on delete restrict,
    file_id       int         not null references storage_files (id) on delete restrict,
    created_at    timestamptz not null default now()
);

create table if not exists report_types
(
    id   int primary key generated always as identity,
    name text not null
);

insert into report_types (name)
values ('Basic');

create table if not exists reports
(
    id           int primary key generated always as identity,
    type         int references report_types (id) on delete restrict,
    period_start date        not null,
    period_end   date        not null,
    created_at   timestamptz not null default now()
);

create table if not exists report_attachments
(
    id         int primary key generated always as identity,
    report_id  int         not null references reports (id) on delete cascade,
    file_id    int         not null references storage_files (id) on delete restrict,
    created_at timestamptz not null default now()
);

create index if not exists idx_users_role on users (role_id);
create index if not exists idx_brigades_status on brigades (status);
create index if not exists idx_subscribers_status on subscribers (status);
create index if not exists idx_passports_subscriber on passports (subscriber_id);
create index if not exists idx_devices_object on devices (object_id);
create index if not exists idx_devices_place_type on devices (place_type);
create index if not exists idx_seals_device on seals (device_id);
create index if not exists idx_contracts_subscriber on contracts (subscriber_id);
create index if not exists idx_contracts_object on contracts (object_id);
create index if not exists idx_tasks_brigade on tasks (brigade_id);
create index if not exists idx_tasks_object on tasks (object_id);
create index if not exists idx_tasks_plan_visit_at on tasks (plan_visit_at);
create index if not exists idx_tasks_status on tasks (status);
create index if not exists idx_inspections_task on inspections (task_id);
create index if not exists idx_inspections_status on inspections (status);
create index if not exists idx_inspections_inspect_at on inspections (inspect_at);
create index if not exists idx_devices_inspection on inspected_devices (inspection_id);
create index if not exists idx_seals_inspection on inspected_seals (inspection_id);
create index if not exists idx_attachments_inspection on inspection_attachments (inspection_id);
create index if not exists idx_attachments_type on inspection_attachments (type);
create index if not exists idx_reports_type on reports (type);
create index if not exists idx_reports_period on reports (period_start, period_end);
create index if not exists idx_attachments_report on report_attachments (report_id);

create or replace function update_updated_at_column()
    returns trigger as
$$
begin
    new.updated_at = now();
    return new;
end;
$$ language plpgsql;

-- Применение триггера ко всем таблицам с updated_at
do
$$
    declare
        table_name text;
    begin
        for table_name in
            select t.table_name
            from information_schema.tables t
                     join information_schema.columns c on t.table_name = c.table_name
            where t.table_schema = 'public'
              and c.column_name = 'updated_at'
            loop
                execute format('create trigger trg_%I_updated_at
                       before update on %I
                       for each row execute function update_updated_at_column()',
                               table_name, table_name);
            end loop;
    end
$$;
