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
);
create index idx_member_desc on member (id desc nulls last); -- desc nulls last : large small null