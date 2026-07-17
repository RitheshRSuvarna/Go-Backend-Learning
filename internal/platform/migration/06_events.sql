-- +goose Up

CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    day_session_id UUID NOT NULL,
    type TEXT NOT NULL,
    ts TIMESTAMPTZ NOT NULL,
    payload_json JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_events_day_session
        FOREIGN KEY (day_session_id)
        REFERENCES day_sessions(id)
        ON DELETE CASCADE
);