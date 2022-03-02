# this is 写経
- https://goswagger.io/tutorial/todo-list.html

# demo

## start server
```sh
make run

go run ./cmd/todo-list-server/main.go --port=3333
2022/02/24 10:26:49 Serving todo list at http://127.0.0.1:3333
```

## get todos
```sh
make get

curl -i 127.0.0.1:3333
HTTP/1.1 200 OK
Content-Type: application/io.goswagger.examples.todo-list.v1+json
Date: Thu, 24 Feb 2022 01:38:34 GMT
Content-Length: 3

[]
```

## add todo
```sh
make add

curl -i localhost:3333 -X POST -H 'Content-Type: application/io.goswagger.examples.todo-list.v1+json' -d "{\"description\":\"message ANDOM\"}"
HTTP/1.1 201 Created
Content-Type: application/io.goswagger.examples.todo-list.v1+json
Date: Thu, 24 Feb 2022 01:39:13 GMT
Content-Length: 39

{"description":"message ANDOM","id":1}
```

## modify todo
```sh
make modify

curl -i localhost:3333/1 -X PUT -H 'Content-Type: application/io.goswagger.examples.todo-list.v1+json' -d '{"description":"go shopping"}'
HTTP/1.1 200 OK
Content-Type: application/io.goswagger.examples.todo-list.v1+json
Date: Thu, 24 Feb 2022 01:39:41 GMT
Content-Length: 37

{"description":"go shopping","id":1}
```



## delete todo
```sh
make delete

curl -i localhost:3333/1 -X DELETE -H 'Content-Type: application/io.goswagger.examples.todo-list.v1+json' 
HTTP/1.1 204 No Content
Date: Thu, 24 Feb 2022 01:40:03 GMT
```

check
```sh
 make get
curl -i 127.0.0.1:3333
HTTP/1.1 200 OK
Content-Type: application/io.goswagger.examples.todo-list.v1+json
Date: Thu, 24 Feb 2022 01:40:24 GMT
Content-Length: 3

[]
```


## not found error
```sh
make modify
curl -i localhost:3333/1 -X PUT -H 'Content-Type: application/io.goswagger.examples.todo-list.v1+json' -d '{"description":"go shopping"}'
HTTP/1.1 404 Not Found
Content-Type: application/io.goswagger.examples.todo-list.v1+json
Date: Thu, 24 Feb 2022 01:40:58 GMT
Content-Length: 41

{"code":404,"message":"not fountd id:1"}
```

# Mock Server
- useing https://github.com/danielgtaylor/apisprout
- onliy first time
    ```sh
    make mock-setup
    ```
- run mock server
    ```sh
    make mock-run
    ```
- check
    ```sh
    curl -i 127.0.0.1:8000
    ```

