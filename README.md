go get -u github.com/go-chi/chi/v5
go get -u github.com/mattn/go-sqlite3

curl -X POST localhost:8080/apply -d 'personal_id=14&name=Super+Mario&amount=500&term=6'