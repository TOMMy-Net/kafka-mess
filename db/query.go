package db

import (
	"context"
	"fmt"
	"os"

	"github.com/TOMMy-Net/kafka-mess/internal/models"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var database *sqlx.DB

func Connect() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	data, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name))
	if err != nil {
		return err
	}
	
	err = migrateBase(data)
	if err != nil {
		return err
	}
	database = data
	return nil
}

func migrateBase(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

	m.Up()
	return nil
}

func WriteMessage(ctx context.Context, message models.Message) (string, error) {
	id := uuid.New().String()
	_, err := database.NamedExecContext(ctx, "INSERT INTO messages (id, message, status) VALUES (:id, :text, :status)", map[string]interface{}{
		"id": id,
		"text": message.Text,
		"status": 0,
	})
	if err != nil {
		return "", err
	}

	return id, nil
}
