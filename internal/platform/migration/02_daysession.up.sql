CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS day_sessions (
    id UUID PRIMARY KEY,
    trip_id UUID NOT NULL,
    date DATE NOT NULL,
    start_time TEXT NOT NULL,
    start_label TEXT NOT NULL,
    start_lat FLOAT8,
    start_lon FLOAT8,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_day_sessions_trip
        FOREIGN KEY (trip_id)
        REFERENCES trips(id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_daysession_created_at ON day_sessions (created_at DESC);
