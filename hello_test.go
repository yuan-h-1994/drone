package main

import (
    "testing"
)

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
	//user, err = storage.FindUser(context.TODO(), 1)
	// ...
	err = storage.CreateUser(context.TODO(), "1")
	if err != nil {
		t.Fatal(err)
	}
}
