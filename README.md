
# API TINGGAL NIKAH








## ENUM POSTGRESQL

 - CREATE TYPE role_status AS ENUM ('admin', 'customer');

## migration
 - change model in folder model
 - if you add new model dont forget create migration file in folder migration
 - load struct from migration file you created
 - uncomment auto migrate gorm in app.go

 

