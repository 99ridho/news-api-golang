# News API

This repository is intended to learn how to implement Clean Architecture in Go & also for assignment purposes.

### How to develop & run

* Go get this repo

    ```
    go get gitlab.com/99ridho/news-api
    ```

* Make the database
* Change the `config.json` file to your own setting
* Run the migration using [`goose`](https://github.com/pressly/goose).
* Install dependency using `dep`

    ```
    dep ensure
    ```

* Then

    ```
    go run main.go
    ```

    at this directory.

### Testing

```
go test ./...
```