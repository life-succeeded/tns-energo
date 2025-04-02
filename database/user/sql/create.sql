insert into users (email, surname, name, patronymic, position, password_hash, refresh_token, refresh_token_expires_at, role_id)
values (:email, :surname, :name, :patronymic, :position, :password_hash, :refresh_token, :refresh_token_expires_at, :role_id)
returning id;