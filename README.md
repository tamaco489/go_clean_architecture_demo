## Go Clean Architecture

### Start API Server
```
$ go version
go version go1.21.3 darwin/arm64
```
```
$ go run main.go

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.11.3
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8080
```


### Exec API Request

* find task by id
    ```
    curl -i -X GET http://localhost:8080/tasks/1 \
        -H "Content-Type: application/json"
    ```

* create task
    ```
    curl -i -X POST http://localhost:8080/tasks \
        -H "Content-Type: application/json" \
        -d '{"title":"New Task title"}'
    ```

### Go Test
```
go test -v -covermode=count ./...
```