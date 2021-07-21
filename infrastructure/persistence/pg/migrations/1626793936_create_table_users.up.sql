create table if not exists users
(
    id         int primary key not null,
    username   varchar(50)     not null,
    last_name  varchar(50)     not null,
    first_name varchar(50)     not null,
    email      varchar(50)     not null,
    phone      varchar(15)     not null
);

create sequence if not exists users_id_seq owned by users.id;