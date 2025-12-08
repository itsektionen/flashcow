package service_test

import (
	"context"
	"errors"
	"slices"
	"testing"

	"github.com/itsektionen/flashcow/backend/internal/model"
	"github.com/itsektionen/flashcow/backend/internal/service"
)

var mockCommittees = []model.Committee{
	{ID: 1, Name: "Qmisk"},
	{ID: 2, Name: "TMEIT"},
	{ID: 3, Name: "ITK"},
}

type mockCommitteeRepo struct {
	ListFunc func(ctx context.Context) ([]model.Committee, error)
	GetFunc  func(ctx context.Context, id int64) (*model.Committee, error)
}

func (m *mockCommitteeRepo) ListCommittees(ctx context.Context) ([]model.Committee, error) {
	return m.ListFunc(ctx)
}

func (m *mockCommitteeRepo) GetCommittee(ctx context.Context, id int64) (*model.Committee, error) {
	return m.GetFunc(ctx, id)
}

func TestCommitteeService_ListCommittees(t *testing.T) {
	ctx := context.Background()

	mockRepo := &mockCommitteeRepo{
		ListFunc: func(ctx context.Context) ([]model.Committee, error) {
			return mockCommittees, nil
		},
	}

	svc := service.NewCommitteeService(mockRepo)

	committees, err := svc.ListCommittees(ctx)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(committees) != 3 {
		t.Fatalf("expected 3 committees, got %d", len(committees))
	}

	if committees[0].Name != "Qmisk" || committees[1].Name != "TMEIT" || committees[2].Name != "ITK" {
		t.Errorf("unexpected committee names: %+v", committees)
	}
}

func TestCommitteeService_GetCommittee(t *testing.T) {
	ctx := context.Background()

	mockRepo := &mockCommitteeRepo{
		GetFunc: func(ctx context.Context, id int64) (*model.Committee, error) {
			idx := slices.IndexFunc(mockCommittees, func(c model.Committee) bool {
				return c.ID == id
			})

			if idx == -1 {
				return nil, errors.New("user not found")
			}

			committee := &mockCommittees[idx]

			return committee, nil
		},
	}

	svc := service.NewCommitteeService(mockRepo)

	// Test existing committee
	c, err := svc.GetCommittee(ctx, 1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if c.Name != "Qmisk" {
		t.Errorf("expected Qmisk, got %s", c.Name)
	}

	// Test non-existing committee
	_, err = svc.GetCommittee(ctx, 4)
	if err == nil {
		t.Fatalf("expected error for non-existing committee")
	}
}
