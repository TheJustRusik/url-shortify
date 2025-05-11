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
    shortcode  VARCHAR(20) NOT NULL,
    ip         VARCHAR(100),
    user_agent TEXT,
    visited_at TIMESTAMP
);

