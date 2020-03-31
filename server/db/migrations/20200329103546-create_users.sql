
-- +migrate Up
CREATE TABLE
IF NOT EXISTS users
(id int
(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
 age int
(11) DEFAULT NULL,
 job varchar
(255) COLLATE utf8_unicode_ci DEFAULT NULL,
 deleted_at datetime DEFAULT NULL,
 created_at datetime DEFAULT NULL,
 updated_at datetime DEFAULT NULL
);
-- +migrate Down
DROP TABLE IF EXISTS users;