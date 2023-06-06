START TRANSACTION;
CREATE TABLE IF NOT EXISTS users(
    id int auto_increment PRIMARY KEY,
    username VARCHAR(255),
    password VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS tasks(
    id int auto_increment PRIMARY KEY,
    user_id int,
    title VARCHAR(255),
    description VARCHAR(255),
    status BOOLEAN,
    created_at TIMESTAMP,
    constraint fk_user_id FOREIGN key (user_id) REFERENCES users(id)
);
COMMIT;