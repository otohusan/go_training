-- +migrate Up
ALTER TABLE users ALTER COLUMN email DROP NOT NULL;
