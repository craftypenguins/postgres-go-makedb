# postgres-go-makedb

Connects to POSTGRES_URL and attempts to create DB_TO_MAKE.

If DB_OWNER and DB_OWNER_PWD are set, it creates that user with that password and sets it as the owner of DB_TO_MAKE. 

These are all environment variables.
