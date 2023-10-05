CREATE TABLE message_queue (
    id SERIAL PRIMARY KEY,
    message TEXT NOT NULL,
    timestamp TIMESTAMPTZ DEFAULT NOW()
);