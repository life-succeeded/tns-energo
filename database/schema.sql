-- Пользователи
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

-- Бригады
create table if not exists brigades
(
    id         int primary key generated always as identity,
    status     int         not null, -- 0 - Idle, 1 - OnTask, 2 - Archived
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

-- Участники бригады
create table if not exists brigade_members
(
    brigade_id   int         not null references brigades (id) on delete cascade,
    inspector_id int         not null references users (id) on delete cascade,
    assigned_at  timestamptz not null default now(),
    primary key (brigade_id, inspector_id)
);

-- Паспорта абонентов
create table if not exists passports
(
    id            int primary key generated always as identity,
    subscriber_id int  not null references subscribers (id) on delete cascade,
    series        text not null check ( series ~ '^\d{4}$' ), -- Серия
    number        text not null check ( number ~ '^\d{6}$' ), -- Номер
    issued_by     text not null,                              -- Кем выдан
    issue_date    date not null                               -- Когда выдан
);

-- Абоненты
create table if not exists subscribers
(
    id             int primary key generated always as identity,
    account_number text        not null unique,                        -- Лицевой счет
    surname        text        not null,
    name           text        not null,
    patronymic     text,
    phone_number   text        not null check ( phone_number ~ '^(\+7|8)\d{10}$' ),
    email          text        not null check ( email ~ '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$' ),
    inn            text        not null check ( inn ~ '^\d{10,12}$' ), -- ИНН
    birth_date     date        not null check ( '1900-01-01' < birth_date and birth_date < current_date ),
    status         int         not null,                               -- 0 - Active, 1 - Violator, 2 - Archived
    created_at     timestamptz not null default now(),
    updated_at     timestamptz not null default now()
);

-- Объекты проверки
create table if not exists inspection_objects
(
    id             int primary key generated always as identity,
    address        text        not null,
    have_automaton bool        not null, -- Наличие коммутационного (вводного) аппарата
    created_at     timestamptz not null default now(),
    updated_at     timestamptz not null default now()
);

-- Приборы учета
create table if not exists devices
(
    id                int primary key generated always as identity,
    object_id         int         not null references inspection_objects (id) on delete cascade,
    type              text        not null,
    number            text        not null,
    place_type        int         not null, -- 0 - Other, 1 - Flat, 2 - StairLanding
    place_description text,                 -- Место установки прибора учета (только для place_type = 0)
    created_at        timestamptz not null default now(),
    updated_at        timestamptz not null default now()
);

-- Пломбы прибора учета
create table if not exists seals
(
    id         int primary key generated always as identity,
    device_id  int         not null references devices (id) on delete cascade,
    number     text        not null,
    place      text        not null, -- Место установки
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

-- Договора абонента
create table if not exists contracts
(
    id            int primary key generated always as identity,
    number        text        not null unique,
    subscriber_id int         not null references subscribers (id) on delete restrict,
    object_id     int         not null references inspection_objects (id) on delete restrict,
    sign_date     date        not null check ( sign_date <= current_date ), -- Дата подписания
    created_at    timestamptz not null default now(),
    updated_at    timestamptz not null default now()
);

-- Задачи
create table if not exists tasks
(
    id            int primary key generated always as identity,
    brigade_id    int references brigades (id),
    object_id     int         not null references inspection_objects (id) on delete cascade,
    plan_visit_at timestamptz,          -- Запланированное время визита
    status        int         not null, -- 0 - Planned, 1 - InWork, 2 - Done
    comment       text,
    started_at    timestamptz,
    finished_at   timestamptz,
    created_at    timestamptz not null default now(),
    updated_at    timestamptz not null default now(),
    check ( started_at is null or finished_at is null or started_at <= finished_at )
);

-- Проверки
create table if not exists inspections
(
    id                        int primary key generated always as identity,
    task_id                   int         not null references tasks (id),
    status                    int         not null, -- 0 - InWork, 1 - Done
    type                      int,                  -- Тип проверки. 0 - Limitation, 1 - Resumption, 2 - Verification, 3 - UnauthorizedConnection
    act_number                text,
    resolution                int,                  -- Результаты проверки. 0 - Limited, 1 - Stopped, 2 - Resumed
    limit_reason              text,                 -- Основание введения ограничения (приостановления) режима потребления. Если NULL, то это неполная оплата
    method                    text,                 -- Способ введения ограничения, приостановления, возобновления режима потребления, номера и место установки пломб
    method_by                 int,                  -- Кем было введено. 0 - Consumer, 1 - Inspector
    reason_type               int,                  -- 0 - NotIntroduced, 1 - ConsumerLimited, 2 - InspectorLimited, 3 - Resumed
    reason_description        text,                 -- Причина невведения ограничения (только для reason_type = 0)
    is_restriction_checked    bool,                 -- Произведена проверка введенного ограничения
    is_violation_detected     bool,                 -- Нарушение потребителем введенного ограничения выявлено
    is_expense_available      bool,                 -- Наличие расхода после введенного ограничения
    violation_description     text,                 -- Иное описание выявленного нарушения/сведения, на основании которых сделан вывод о нарушении
    is_unauthorized_consumers bool,                 -- Самовольное подключение энергопринимающих устройств Потребителя к электрическим сетям
    unauthorized_description  text,                 -- Описание места и способа самовольного подключения к электрическим сетям
    unauthorized_explanation  text,                 -- Объяснение лица, допустившего самовольное подключение к электрическим сетям
    inspect_at                timestamptz,          -- Дата проверки
    energy_action_at          timestamptz,          -- Время действия над подачей электроэнергии
    created_at                timestamptz not null default now(),
    updated_at                timestamptz not null default now()
);

-- Проверенные приборы учета
create table if not exists inspected_devices
(
    id            int primary key generated always as identity,
    device_id     int            not null references devices (id) on delete cascade,
    inspection_id int            not null references inspections (id) on delete cascade,
    value         numeric(15, 2) not null, -- Текущее показание
    consumption   numeric(15, 2) not null  -- Расход электрической энергии кВтч
);

-- Проверенные пломбы
create table if not exists inspected_seals
(
    id            int primary key generated always as identity,
    seal_id       int  not null references seals (id) on delete cascade,
    inspection_id int  not null references inspections (id) on delete cascade,
    is_broken     bool not null -- Сорвана ли пломба
);

-- Файлы
create table if not exists storage_files
(
    id        int primary key generated always as identity,
    file_name text not null,
    file_size int  not null,
    bucket    text not null,
    url       text not null -- Ссылка на скачивание
);

-- Приложения к проверкам
create table if not exists inspection_attachments
(
    id            int primary key generated always as identity,
    inspection_id int         not null references inspections (id) on delete cascade,
    type          int         not null, -- 0 - DevicePhoto, 1 - SealPhoto, 2 - Act
    file_id       int         not null references storage_files (id) on delete restrict,
    created_at    timestamptz not null default now()
);

-- Отчеты
create table if not exists reports
(
    id         int primary key generated always as identity,
    type       int         not null, -- 0 - Daily, 1 - Weekly, 2 - Monthly
    file_id    int         not null references storage_files (id) on delete restrict,
    start_date date        not null,
    end_date   date        not null,
    created_at timestamptz not null default now(),
    check ( start_date <= end_date )
);

create index if not exists idx_users_role on users (role_id);
create index if not exists idx_brigades_status on brigades (status);
create index if not exists idx_subscribers_status on subscribers (status);
create index if not exists idx_devices_object on devices (object_id);
create index if not exists idx_seals_device on seals (device_id);
create index if not exists idx_tasks_brigade on tasks (brigade_id);
create index if not exists idx_tasks_object on tasks (object_id);
create index if not exists idx_tasks_status on tasks (status);
create index if not exists idx_inspections_task on inspections (task_id);
create index if not exists idx_inspections_status on inspections (status);
create index if not exists idx_inspection_attachments_inspection on inspection_attachments (inspection_id);
create index if not exists idx_inspected_devices_inspection on inspected_devices (inspection_id);
create index if not exists idx_inspected_seals_inspection on inspected_seals (inspection_id);
create index if not exists idx_passports_subscriber on passports (subscriber_id);
create index if not exists idx_contracts_subscriber on contracts (subscriber_id);
create index if not exists idx_inspected_devices_device on inspected_devices (device_id);
create index if not exists idx_inspected_seals_seal on inspected_seals (seal_id);
create index if not exists idx_reports_dates on reports (start_date, end_date);
create index if not exists idx_tasks_plan_visit_at on tasks (plan_visit_at);
create index if not exists idx_inspections_inspect_at on inspections (inspect_at);

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
                execute format('drop trigger if exists trg_%I_updated_at on %I', table_name, table_name);
                execute format('create trigger trg_%I_updated_at
                       before update on %I
                       for each row execute function update_updated_at_column()',
                               table_name, table_name);
            end loop;
    end
$$;
