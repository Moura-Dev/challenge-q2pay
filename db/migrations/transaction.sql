create type status as enum ('OPEN', 'DONE', 'CANCELED');

create table if not exists transactions (
    id serial,
    value decimal(10,4) NOT NULL, CONSTRAINT value_positive CHECK (value > 0),
    payer int NOT NULL,
    payee int NOT NULL,
    status status not null,
    created_at timestamp default NOW(),
    updated_at timestamp default NOW(),
    PRIMARY KEY (id),
    FOREIGN KEY (payer) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (payee) REFERENCES users(id) ON DELETE CASCADE
);

