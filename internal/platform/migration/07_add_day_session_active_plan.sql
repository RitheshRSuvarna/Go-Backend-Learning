-- +goose Up

ALTER TABLE day_sessions
ADD COLUMN active_plan_version_id UUID
REFERENCES plan_versions(id);