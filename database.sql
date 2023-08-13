/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

/** This is test table. Remove this table and replace with your own tables. */
CREATE TABLE users (
	id serial PRIMARY KEY,
	phone_number VARCHAR ( 20 ) UNIQUE NOT NULL,
  full_name VARCHAR ( 60 ) NOT NULL,
  password TEXT,
  success_login_count int DEFAULT 0
);

INSERT INTO users (phone_number, full_name, password) VALUES ('+6218426881', 'Erma', 'hashed_and_salted_password');
