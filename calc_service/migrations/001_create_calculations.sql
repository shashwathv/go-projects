CREATE TABLE IF NOT EXISTS calculations (
    id SERIAL PRIMARY KEY,
    operation TEXT NOT NULL,
    operands INTEGER[] NOT NULL,
    result INTEGER NOT NULL,
    request_id TEXT NOT NULL,
    user_id TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);
