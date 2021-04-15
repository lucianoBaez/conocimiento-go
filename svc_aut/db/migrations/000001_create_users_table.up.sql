CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   name     VARCHAR(50) NOT NULL,
   username VARCHAR (50) UNIQUE NOT NULL,
   password VARCHAR (256) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NULL,
   deleted_at TIMESTAMP NULL
);

insert into users (name, username, password, email, created_at)
values ('admin','admin', '$2a$10$NP5DLE/laTBL3wFJUKCu9.jrH9SYVC6sMpQvsD8XRZU7Km2IXlsvK', 'admin@toucan.com', now());