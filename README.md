CRUD Приложение

Стэк

go 1.22.1

postgres



Для postgres можно использовать Docker

docker run -d --name ninja-db -e POSTGRES_PASSWORD=qwerty123 -v ${HOME}/pgdata/:/var/lib/postgresql/data -p 5432:5432 postgres
