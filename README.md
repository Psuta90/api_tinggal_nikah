
# API TINGGAL NIKAH








## ENUM POSTGRESQL

 - CREATE TYPE role_status AS ENUM ('admin', 'customer');

## migration
 - change model in folder model
 - if you add new model dont forget create migration file in folder migration
 - load struct from migration file you created
 - uncomment auto migrate gorm in app.go

## setting rbac
 - first migrate
 - and insert policy in table casbin_rule
 - for now rbac just by group route and role
 - template query : INSERT INTO public.casbin_rule
(id, ptype, v0, v1, v2, v3, v4, v5)
VALUES(2, 'p', 'customer', '/api/v1/customer/*', 'POST', NULL, NULL, NULL);


 

