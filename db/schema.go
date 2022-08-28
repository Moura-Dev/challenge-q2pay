package db

const Schema = `
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

create table if not exists transactions (
    id serial,
    value decimal(10,4) NOT NULL, CONSTRAINT value_positive CHECK (value > 0),
    payer int NOT NULL,
    payee int NOT NULL,
    created_at timestamp default NOW(),
    PRIMARY KEY (id),
    FOREIGN KEY (payer) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (payee) REFERENCES users(id) ON DELETE CASCADE
);
create table if not exists wallets (
    id serial,
    balance decimal(10,4) default 0 CONSTRAINT balance_positive CHECK (balance >= 0) ,
    user_id int NOT NULL ,
    version int default 0,
    Primary KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);`
