package main

import (
    "testing"
    "fmt"
    "context"
)
type DataStore interface {
	FindUser(ctx context.Context, id int) (*User, error)
	CreateUser(ctx context.Context, name string) error
}

type User struct {
	id   int
	name string
}

type MockUserDataStore struct{}

func NewMockDataStore() DataStore {
	return &MockUserDataStore{}
}

func (m *MockUserDataStore) FindUser(ctx context.Context, id int) (*User, error) {
	return &User{
		id:   1,
		name: "test user",
	}, nil
}

func (m *MockUserDataStore) CreateUser(ctx context.Context, name string) error {
	fmt.Println("create user. dummy")
	return nil
}

func TestHello(t *testing.T) {
    if "Hello World" != getHello() {
        t.Fatal("failed test")
    }
    storage := NewMockDataStore()
	user, err := storage.FindUser(context.TODO(), 1)
	if user.id != 1 {
		t.Fatal(err)
	}
	err1 := storage.CreateUser(context.TODO(), "1")
	if err1 != nil {
		t.Fatal(err)
	}
}
