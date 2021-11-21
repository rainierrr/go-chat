
-- +migrate Up
CREATE TABLE IF NOT EXISTS channels (
    id INTEGER UNSIGNED PRIMARY KEY  AUTO_INCREMENT,
    name VARCHAR(128) NOT NULL,
    owner INTEGER UNSIGNED NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) DEFAULT CHARSET=utf8mb4;
-- +migrate Down
DROP TABLE IF EXISTS channels;
