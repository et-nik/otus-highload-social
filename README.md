# Otus Highload Social Network

Project for the OTUS Highload Architect course

## Run Test Database in Docker

```
docker run --rm --name mysql -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=social -p 13306:3306 mysql
```

## Run migrations

```
cd migrations
goose mysql "root:password@/social?parseTime=true" up
```


## Run application

```
PORT=8080 \
DATABASE=root:password@(localhost:13306)/social \
./social
```
