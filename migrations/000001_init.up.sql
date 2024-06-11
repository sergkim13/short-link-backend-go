CREATE TABLE links
(
    id          SERIAL PRIMARY KEY,
    original    VARCHAR(255) NOT NULL,
    short       VARCHAR(255) NOT NULL,
    clicks      INT NOT NULL DEFAULT 0,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);