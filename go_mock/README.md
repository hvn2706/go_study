# Demo go mock project

## introduce

-   Api web service, using gin framework
-   Demo:
    +create:
    curl http://localhost:8080/albums \
     --include \
     --header "Content-Type: application/json" \
     --request "POST" \
     --data '{"title": "The Modern","artist": "hvn","price": 50.12}'

    +edit:
    curl http://localhost:8080/albums/edit \
     --include \
     --header "Content-Type: application/json" \
     --request "PUT" \
     --data '{"id": 4,"title": "The Modern","artist": "hvn","price": 50.12}'

    +get:
    curl http://localhost:8080/albums/4 \
     --include \
     --header "Content-Type: application/json" \
     --request "GET"

    +delete:
    curl http://localhost:8080/albums/4 \
     --include \
     --header "Content-Type: application/json" \
     --request "DELETE"

## How to run

-   Presequisite:
    -   go@1.18 and above
    -   python3.6 and above
    -   mysql server
-   Create database by running sql script in database/create-table.sql
-   Open terminal, go to main module directory, run command:
    > go run main.go
-   Locate to Client demo, run command:
    > python3 -m http.server 3000
-   Open browser, go to http://localhost:3000
