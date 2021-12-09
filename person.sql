--DROP DATABASE person;

CREATE DATABASE person;

DROP TABLE IF EXISTS article;
CREATE TABLE person (
  id serial,
  email varchar(150) NOT NULL,
  phone varchar(12) NOT NULL,
  firstname varchar(150) DEFAULT NULL,
  lastname varchar(150) DEFAULT NULL,
  PRIMARY KEY (id)
)

