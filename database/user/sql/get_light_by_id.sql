select id,
       email,
       surname,
       name,
       patronymic,
       position
from users
where id = $1;