CREATE TABLE IF NOT EXISTS short_urls
(
    id           SERIAL PRIMARY KEY,
    shortcode    VARCHAR(20) UNIQUE NOT NULL,
    original_url TEXT               NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE visits
(
    id         SERIAL PRIMARY KEY,
    short_url  TEXT NOT NULL,
    ip         TEXT,
    user_agent TEXT,
    referer    TEXT,
    timestamp  TIMESTAMP
);
