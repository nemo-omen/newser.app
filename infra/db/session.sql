-- database: ../../data/newser.sqlite
CREATE TABLE IF NOT EXISTS
    sessions (
        token CHAR(43) PRIMARY KEY,
        data BLOB NOT NULL,
        expiry TIMESTAMP (6) NOT NULL
    );

DROP TABLE IF EXISTS sessions;
