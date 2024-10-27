-- Create table for CPU usage
CREATE TABLE cpu_usage (
                           id SERIAL PRIMARY KEY,
                           process_name VARCHAR(255) NOT NULL,
                           usage FLOAT NOT NULL,
                           timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);