CREATE TABLE users
(
    id   serial       not null unique,
    name varchar(255) not null,
    age  int          not null,
    sex  varchar(255) not null
);