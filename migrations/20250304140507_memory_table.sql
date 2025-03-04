-- +goose Up
-- +goose StatementBegin
CREATE TABLE memory_usage (
    process_name VARCHAR(255) NOT NULL,
    usage FLOAT NOT NULL,
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE memory_usage;
-- +goose StatementEnd
