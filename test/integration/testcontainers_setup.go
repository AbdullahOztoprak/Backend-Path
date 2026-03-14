//go:build integration

package integration

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	postgresContainer testcontainers.Container
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	// Set up PostgreSQL container
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpassword",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	var err error
	postgresContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Failed to start PostgreSQL container: %s", err)
	}

	// Run tests
	code := m.Run()

	// Clean up
	if err := postgresContainer.Terminate(ctx); err != nil {
		log.Fatalf("Failed to terminate PostgreSQL container: %s", err)
	}

	os.Exit(code)
}

func GetPostgresURI() string {
	host, err := postgresContainer.Host(context.Background())
	if err != nil {
		log.Fatalf("Failed to get container host: %s", err)
	}

	port, err := postgresContainer.MappedPort(context.Background(), "5432")
	if err != nil {
		log.Fatalf("Failed to get mapped port: %s", err)
	}

	return fmt.Sprintf("host=%s port=%s user=testuser password=testpassword dbname=testdb sslmode=disable", host, port.Port())
}