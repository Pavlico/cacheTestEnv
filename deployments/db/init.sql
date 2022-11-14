CREATE DATABASE IF NOT EXISTS userdb;
USE userdb;
CREATE TABLE IF NOT EXISTS users (
  id int(11) NOT NULL AUTO_INCREMENT,
  info varchar(255),
  CONSTRAINT users PRIMARY KEY (id)
);
INSERT INTO users (id, info) VALUES 
(1, "TEST"),
(2, "TEST"),
(3, "TEST"),
(4, "TEST"),
(5, "TEST"),
(6, "TEST")