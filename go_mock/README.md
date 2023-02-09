# Demo go mock project

    - Api web service, using gin framework
    - Demo:
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
