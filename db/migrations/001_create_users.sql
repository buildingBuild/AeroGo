CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users(
    id uuid PRIMARY KEY DEFAULT gen_random_uuid()
    number TEXT UNIQUE NOT NULL
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP
)
