DROP DATABASE IF EXISTS restapi;

CREATE DATABASE restapi;

-- connect manually in psql
-- \c restapi
CREATE TABLE
    users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        password_hash TEXT NOT NULL,
        last_login TIMESTAMPTZ,
        is_admin BOOLEAN DEFAULT FALSE,
        is_active BOOLEAN DEFAULT TRUE,
        created_at TIMESTAMPTZ DEFAULT NOW ()
    );

INSERT INTO
    users (
        username,
        password_hash,
        last_login,
        is_admin,
        is_active
    )
VALUES
    ('admin', 'admin', NOW (), TRUE, TRUE);

-------------------- How To run? -------------------
-- psql -U postgres -f schema.sql