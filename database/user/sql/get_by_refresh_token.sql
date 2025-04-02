select id,
       email,
       surname,
       name,
       patronymic,
       position,
       password_hash,
       refresh_token,
       refresh_token_expires_at,
       role_id
from users
where refresh_token = $1;