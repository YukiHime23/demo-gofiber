create table users
(
    id bigserial constraint users_pk primary key,
    name varchar(255),
    email varchar(255),
    password varchar(255)
)
