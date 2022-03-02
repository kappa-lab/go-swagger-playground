swag-valid:
	swagger validate ./swagger.yml
swag-gen:	
	swagger generate server -A todo-list -f ./swagger.yml

swag-gen-client:	
	swagger generate client -A todo-list -f ./swagger.yml
	
oas-valid:
	openapi-generator validate -i openapi.yml

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

test:
	go test -v ./client 	

mock-setup:
	docker pull danielgtaylor/apisprout

mock-run:
	docker run -p 8000:8000 -v `pwd`/openapi.yml:/api.yaml danielgtaylor/apisprout /api.yaml
