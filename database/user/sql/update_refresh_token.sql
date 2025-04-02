update users
set refresh_token            = $2,
    refresh_token_expires_at = $3
where id = $1;