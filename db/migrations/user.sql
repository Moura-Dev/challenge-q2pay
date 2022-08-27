CREATE TABLE IF NOT EXISTS users (
    id SERIAL,
    full_name varchar(255) NOT NULL,
    cpfcnpj varchar(255) unique NOT NULL,
    email varchar(255) unique NOT NULL,
    login varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    active boolean default True NOT NULL,
    created_at timestamp default NOW(),
    updated_at timestamp default NOW(),
    PRIMARY KEY (id)
);

insert into users (full_name, cpfcnpj, email, login, password) values ('John Doe', '12345678944', 'jr@jrr', 'jrjrjr', 'testestae');

select * from users;