-- +goose Up
CREATE TABLE plan_versions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    day_session_id UUID NOT NULL,
    version INTEGER NOT NULL,
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_plan_versions_day_session
        FOREIGN KEY (day_session_id)
        REFERENCES day_sessions(id)
        ON DELETE CASCADE,

    CONSTRAINT uq_plan_versions_day_session_version
        UNIQUE (day_session_id, version)
);

CREATE TABLE plan_stops (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    plan_version_id UUID NOT NULL,
    position INTEGER NOT NULL,
    title TEXT NOT NULL,
    category_label TEXT NOT NULL,
    image_url TEXT,
    planned_arrival TIMESTAMPTZ NOT NULL,
    planned_departure TIMESTAMPTZ NOT NULL,
    travel_minutes INTEGER NOT NULL,
    stay_minutes INTEGER NOT NULL,
    busy_risk_label TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_plan_stops_plan_version
        FOREIGN KEY (plan_version_id)
        REFERENCES plan_versions(id)
        ON DELETE CASCADE,

    CONSTRAINT uq_plan_stops_position
        UNIQUE (plan_version_id, position)
);

-- -- +goose Down

-- DROP TABLE IF EXISTS plan_stops;
-- DROP TABLE IF EXISTS plan_versions;