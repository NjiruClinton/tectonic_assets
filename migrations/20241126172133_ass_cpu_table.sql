-- +goose Up
-- +goose StatementBegin
CREATE TABLE cpu_usage (
                           id SERIAL PRIMARY KEY,
                           process_name VARCHAR(255) NOT NULL,
                           usage FLOAT NOT NULL,
                           timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cpu_usage;
-- +goose StatementEnd
