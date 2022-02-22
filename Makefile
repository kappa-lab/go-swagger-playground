swag-valid:
	swagger validate ./swagger.yml
swag-gen:	
	swagger generate server -A todo-list -f ./swagger.yml