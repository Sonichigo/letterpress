1. To start the application use ```docker-compose up --build```
2. Start the DB Migraton :-
    - a. ```export PGURL="postgres://letterpress:letterpress_secrets@localhost:5432/letterpress_db?sslmode=disable"```
    - b. ```migrate -database $PGURL -path db/migrations/ up```
