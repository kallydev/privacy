CREATE TABLE IF NOT EXISTS qq
(
    id           BIGINT,
    qq_number    BIGINT,
    phone_number INT
);

CREATE TABLE IF NOT EXISTS jd
(
    id           BIGINT,
    name         TEXT,
    nickname     TEXT,
    password     TEXT,
    email        TEXT,
    id_number    TEXT,
    phone_number INT
);

CREATE TABLE IF NOT EXISTS sf
(
    id           BIGINT,
    name         TEXT,
    phone_number INT,
    address      TEXT
);
