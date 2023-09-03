
# API TINGGAL NIKAH








## ENUM POSTGRESQL

 - CREATE TYPE role_status AS ENUM ('admin', 'customer');
 - CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

## migration
 - change env on root file DATABASE_URL with your database
 - run go run migration.go

## setting rbac
 - first migrate
 - and insert policy in table casbin_rule
 - for now rbac just by group route and role
 - template query : INSERT INTO public.casbin_rule
(id, ptype, v0, v1, v2, v3, v4, v5)
VALUES(2, 'p', 'customer', '/api/v1/customer/*', 'POST', NULL, NULL, NULL);


 

