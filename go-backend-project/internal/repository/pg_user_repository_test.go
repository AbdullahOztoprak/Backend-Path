package repository_test

import (
	"testing"
	"github.com/AbdullahOztoprak/go-backend-project/internal/models"
	"github.com/AbdullahOztoprak/go-backend-project/internal/repository"
)

type dummyConn struct{}

func (d *dummyConn) Exec(ctx interface{}, sql string, args ...interface{}) (interface{}, error) { return nil, nil }
func (d *dummyConn) QueryRow(ctx interface{}, sql string, args ...interface{}) *models.User { return &models.User{} }
func (d *dummyConn) Query(ctx interface{}, sql string, args ...interface{}) ([]*models.User, error) { return []*models.User{}, nil }

// func TestPGUserRepositoryCreate(t *testing.T) {
// 	repo := repository.NewPGUserRepository(nil)
// 	user := &models.User{
// 		Username:     "testuser",
// 		Email:        "test@example.com",
// 		PasswordHash: "hashedpassword",
// 		Role:         "user",
// 	}
// 	err := repo.Create(user)
// 	if err != nil {
// 		t.Errorf("expected nil error, got %v", err)
// 	}
// }
