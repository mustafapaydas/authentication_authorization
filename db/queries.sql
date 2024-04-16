CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists tbl_user
(
    id           serial primary key,
    user_id      uuid         not null default uuid_generate_v4() unique,
    created_date TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_by   uuid,
    updated_date TIMESTAMP,
    username     varchar(30)  not null unique,
    first_name   varchar(100) not null,
    last_name    varchar(50)  not null,
    email        varchar(100) NOT null unique,
    full_name    varchar(100),
    password     text         not null,
    phone_number varchar(20) unique,
    is_verify bool default false
);

create table if not exists tbl_role
(
    id           serial primary key,
    role_id      uuid         not null default uuid_generate_v4() unique,
    created_date TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_by   uuid,
    updated_date TIMESTAMP,
    name         varchar(100) NOT null unique
);

create table if not exists tbl_user_role_relation
(
    user_id int,
    role_id int,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES tbl_user (id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES tbl_role (id) ON DELETE CASCADE
);

CREATE OR REPLACE FUNCTION fnc_add_role()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO tbl_user_role_relation(user_id, role_id)
    VALUES (NEW.id, (select id from tbl_role where name = 'GENERAL_USER'));
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE if not exists TRIGGER trg_add_role
AFTER INSERT ON tbl_user
FOR EACH ROW
EXECUTE FUNCTION fnc_add_role();

