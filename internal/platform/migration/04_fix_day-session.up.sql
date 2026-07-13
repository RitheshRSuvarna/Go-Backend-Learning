-- +goose Up
ALTER TABLE day_sessions
ALTER COLUMN id
SET DEFAULT gen_random_uuid();