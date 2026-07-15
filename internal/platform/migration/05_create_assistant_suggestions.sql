-- +goose Up

CREATE TABLE assistant_suggestions (
    id UUID PRIMARY KEY,
    day_session_id UUID NOT NULL,
    message TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_assistant_suggestions_day_session
        FOREIGN KEY (day_session_id)
        REFERENCES day_sessions(id)
        ON DELETE CASCADE,

    CONSTRAINT assistant_suggestions_status_check
        CHECK (status IN ('pending', 'accepted', 'snoozed'))
);