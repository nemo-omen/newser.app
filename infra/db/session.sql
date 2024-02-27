-- database: ../../data/newser.sqlite
CREATE TABLE IF NOT EXISTS
    sessions (
        token TEXT PRIMARY KEY,
        data BLOB NOT NULL,
        expiry REAL NOT NULL
    );

DROP TABLE IF EXISTS sessions;
