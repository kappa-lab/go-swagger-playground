swag-valid:
	swagger validate ./swagger.yml
swag-gen:	
	swagger generate server -A todo-list -f ./swagger.yml

run:
	go run ./cmd/todo-list-server/main.go --port=3333