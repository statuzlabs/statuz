CREATE TABLE monitors (
    id               TEXT PRIMARY KEY,
    name             TEXT NOT NULL,
    type             TEXT NOT NULL,
    url              TEXT,
    interval_sec     INT NOT NULL,
    degraded_thresh_ms INT,
    enabled          BOOLEAN DEFAULT true,
    created_at       TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
