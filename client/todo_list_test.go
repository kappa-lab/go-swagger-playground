package client

import (
	"log"
	"testing"
	"time"

	"github.com/go-openapi/swag"
	"github.com/kappa-lab/go-swagger-playground/client/todos"
	"github.com/kappa-lab/go-swagger-playground/models"
)

var tc = TransportConfig{
	"localhost:3333",
	DefaultBasePath,
	DefaultSchemes,
}

func TestHoge(t *testing.T) {
	got := 1 + 1
	want := 2
	if got != want {
		t.Errorf("Max(1, 2) == %d, want %d", got, want)
	}
}

func TestAddOne(t *testing.T) {
	d := "test1"
	i := models.Item{
		Completed:   false,
		Description: &d,
	}
	p := todos.NewAddOneParams().WithBody(&i)
	resp, err := NewHTTPClientWithConfig(nil, &tc).Todos.AddOne(p.WithTimeout(10 * time.Second))
	if err != nil {
		log.Fatal(err)
	}

	if resp.Payload.Completed == true {
		log.Fatal("Completed must false")
	}

	if *resp.Payload.Description != d {
		log.Fatal("Description not same")
	}

	log.Println(&d)
}

func TestFindTodos(t *testing.T) {
	p := todos.NewFindTodosParams()
	p.Since = swag.Int64(0)
	resp, err := NewHTTPClientWithConfig(nil, &tc).Todos.FindTodos(p.WithTimeout(10 * time.Second))
	if err != nil {
		log.Fatal(err)
	}
	if len(resp.Payload) <= 0 {
		log.Fatal("No records")
	}
	for _, p := range resp.Payload {
		log.Println(p.ID, *p.Description, p.Completed)
	}
}
