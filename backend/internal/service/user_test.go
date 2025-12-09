package service_test

import (
	"context"
	"errors"
	"slices"
	"testing"

	"github.com/itsektionen/flashcow/backend/internal/model"
	"github.com/itsektionen/flashcow/backend/internal/service"
)

var mockUsers = []model.User{
	{ID: 1, FullName: "John Doe", ChapterEmailAddress: "johndoe@kth.it"},
	{ID: 2, FullName: "Jane Smith", ChapterEmailAddress: "janesmith@kth.it"},
	{ID: 3, FullName: "Alice Johnson", ChapterEmailAddress: "alicejohnson@kth.it"},
}

type mockUserRepo struct {
	ListFunc   func(ctx context.Context) ([]model.User, error)
	GetFunc    func(ctx context.Context, id int64) (*model.User, error)
	CreateFunc func(ctx context.Context, user *model.User) error
}

func (m *mockUserRepo) ListUsers(ctx context.Context) ([]model.User, error) {
	return m.ListFunc(ctx)
}

func (m *mockUserRepo) GetUser(ctx context.Context, id int64) (*model.User, error) {
	return m.GetFunc(ctx, id)
}

func (m *mockUserRepo) CreateUser(ctx context.Context, user *model.User) error {
	return m.CreateFunc(ctx, user)
}

func TestUserService_ListUsers(t *testing.T) {
	ctx := context.Background()

	mockRepo := &mockUserRepo{
		ListFunc: func(ctx context.Context) ([]model.User, error) {
			return mockUsers, nil
		},
	}

	svc := service.NewUserService(mockRepo)

	// Test all users
	users, err := svc.ListUsers(ctx)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test out of range
	if len(users) != 3 {
		t.Fatalf("expected 3 users, got %d", len(users))
	}

	// Test all users
	if users[0].FullName != "John Doe" || users[1].FullName != "Jane Smith" || users[2].FullName != "Alice Johnson" {
		t.Errorf("unexpected user names: %+v", users)
	}
}

func TestUserService_GetUser(t *testing.T) {
	ctx := context.Background()

	mockRepo := &mockUserRepo{
		GetFunc: func(ctx context.Context, id int64) (*model.User, error) {
			idx := slices.IndexFunc(mockUsers, func(c model.User) bool {
				return c.ID == id
			})

			if idx == -1 {
				return nil, errors.New("user not found")
			}

			user := &mockUsers[idx]

			return user, nil
		},
	}

	svc := service.NewUserService(mockRepo)

	// Test existing user
	u, err := svc.GetUser(ctx, 1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if u.FullName != "John Doe" {
		t.Errorf("expected John Doe, got %s", u.FullName)
	}

	// Test non-existing user
	_, err = svc.GetUser(ctx, 5)
	if err == nil {
		t.Fatalf("expected error for non-existing user")
	}
}

func TestUserService_CreateUser(t *testing.T) {
	ctx := context.Background()

	mockRepo := &mockUserRepo{
		CreateFunc: func(ctx context.Context, user *model.User) error {
			if user.ID == 0 {
				return errors.New("user ID cannot be 0")
			}

			mockUsers = append(mockUsers, *user)

			return nil
		},
		ListFunc: func(ctx context.Context) ([]model.User, error) {
			return mockUsers, nil
		},
	}

	svc := service.NewUserService(mockRepo)

	// Test creating a user
	user := &model.User{
		ID:                  4,
		FullName:            "Bob Smith",
		ChapterEmailAddress: "bob@example.com",
	}

	newUser, err := svc.CreateUser(ctx, *user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test user was created
	users, err := svc.ListUsers(ctx)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(users) != 4 {
		t.Fatalf("expected 4 users, got %d", len(users))
	}

	found := false
	for _, u := range users {
		if u.ID == newUser.ID {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("expected user %v in list", newUser)
	}
}
