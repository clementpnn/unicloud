CREATE TABLE IF NOT EXISTS links (
    id VARCHAR(36) PRIMARY KEY,
    long_url TEXT NOT NULL,
    short_url VARCHAR(10) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_short_url ON links(short_url);
CREATE INDEX idx_long_url ON links(long_url); 