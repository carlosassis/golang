docker run --name postgres-store \
-e POSTGRES_PASSWORD=mysecretpassword \
-e POSTGRES_DB=storedb \
-p 54320:5432 \
-d postgres 