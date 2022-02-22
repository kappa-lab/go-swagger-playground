swag-valid:
	swagger validate ./swagger.yml
swag-gen:	
	swagger generate server -A todo-list -f ./swagger.yml

run:
	go run ./cmd/todo-list-server/main.go --port=3333

get:
	curl -i 127.0.0.1:3333

add:
	curl -i localhost:3333 -X POST -H 'Content-Type: application/io.goswagger.examples.todo-list.v1+json' -d "{\"description\":\"message $RANDOM\"}"

modify:
	curl -i localhost:3333/1 -X PUT -H 'Content-Type: application/io.goswagger.examples.todo-list.v1+json' -d '{"description":"go shopping"}'

delete:
	curl -i localhost:3333/1 -X DELETE -H 'Content-Type: application/io.goswagger.examples.todo-list.v1+json' 

