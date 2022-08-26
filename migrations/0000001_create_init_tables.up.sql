/*
postgreSQL
*********************************************************************
*/

-- brands table
CREATE TABLE IF NOT EXISTS member
(
    id              SERIAL PRIMARY KEY,
    account         text UNIQUE NOT NULL,
    password        text NOT NULL,
    permission      integer NOT NULL,
    name            text  NOT NULL,
    email           text UNIQUE,
    phone           text UNIQUE,
    address         text ,
    is_alive        boolean NOT NULL DEFAULT true,
    create_at       timestamp with time zone DEFAULT now(),
    update_at       timestamp without time zone,
);
create index idx_member_desc on member (id desc nulls last); -- desc nulls last : large small null

INSERT INTO member(
	account, password, permission, name, email, phone, address)
	VALUES ('admin', '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918', 1, 'admin', 'admin@gmail.com', 0912345678, '台北市內湖陽光街1段2巷33號');
    VALUES ('firm', '6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b', 2, '屋馬', 'wuma@gmail.com', 0912345675, '台北市內湖陽光街1段2巷33號');