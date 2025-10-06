CREATE DATABASE belanjakuy;
ERROR:  database "belanjakuy" already exists
postgres=# DROP DATABASE belanjakuy;
DROP DATABASE
postgres=# CREATE DATABASE belanjakuy;
CREATE DATABASE
postgres=# \c belanjakuy;
You are now connected to database "belanjakuy" as user "postgres".
belanjakuy=# CREATE TABLE customers (customer_id SERIAL PRIMARY KEY, email varchar(100) UNIQUE, phone varchar(100));
CREATE TABLE
belanjakuy=# \dt
           List of relations
 Schema |   Name    | Type  |  Owner
--------+-----------+-------+----------
 public | customers | table | postgres
(1 row)


belanjakuy=# INSERT INTO customers (email, phone) VALUES ('sihab@mail.com', '089123456789'), ('sihabdummy@mail.com', '089987654321');
INSERT 0 2
belanjakuy=# INSERT INTO customers (email, phone) VALUES ('sihab2@mail.com', '089
123456789'), ('sihabdummy1@mail.com', '089987654321');
INSERT 0 2
belanjakuy=# SELECT * FROM customers;
 customer_id |        email         |    phone
-------------+----------------------+--------------
           1 | sihab@mail.com       | 089123456789
           2 | sihabdummy@mail.com  | 089987654321
           3 | sihab2@mail.com      | 089123456789
           4 | sihabdummy1@mail.com | 089987654321
(4 rows)


belanjakuy=# INSERT INTO customers (email, phone) VALUES ('sihab3@mail.com', '089
123456789'), ('sihabdummy2@mail.com', '089987654321');