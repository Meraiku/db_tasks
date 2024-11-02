-- +goose Up
CREATE INDEX IF NOT EXISTS idx_last_name on employees(last_name);

-- +goose Down
DROP INDEX IF EXISTS idx_last_name;
