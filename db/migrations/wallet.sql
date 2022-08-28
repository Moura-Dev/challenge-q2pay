create table if not exists wallets (
    id serial,
    balance decimal(10,4) default 0 CONSTRAINT balance_positive CHECK (balance >= 0) ,
    user_id int NOT NULL ,
    version int default 0,
    Primary KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
