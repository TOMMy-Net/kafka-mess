package db

import (
	"context"
	"errors"
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

type Database struct {
	*sqlx.DB
}

var (
	ErrBaseWrite  = errors.New("error with data write")
	ErrWithSelect = errors.New("error with get data")
)

func Connect() (*Database, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	ssl := os.Getenv("SSL_MODE")

	data, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, name, ssl))
	if err != nil {
		return &Database{}, err
	}

	err = migrateBase(data)
	if err != nil {
		return &Database{}, err
	}
	return &Database{data}, nil
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

func (d *Database) WriteMessage(ctx context.Context, message models.Message) (uuid.UUID, error) {
	id := uuid.New()
	_, err := d.NamedExecContext(ctx, "INSERT INTO messages (uid, message, status) VALUES (:id, :text, :status)", map[string]interface{}{
		"id":     id,
		"text":   message.Text,
		"status": 0,
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (d *Database) UpdateMessageStatus(ctx context.Context, uid string, status int) error {
	_, err := d.NamedExecContext(ctx, "UPDATE messages SET status = :status WHERE uid = :id", map[string]any{
		"id":     uid,
		"status": status,
	})
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) UnSendMessages(ctx context.Context) ([]models.Message, error) {
	var mess []models.Message

	err := d.SelectContext(ctx, &mess, "SELECT * FROM messages WHERE status = 0")
	if err != nil {
		return []models.Message{}, err
	}
	return mess, nil
}

func (d *Database) TotalMessages(ctx context.Context) ([]models.Message, error) {
	var count []models.Message
	err := d.SelectContext(ctx, &count, "SELECT * FROM messages")
	if err != nil {
		return []models.Message{}, err
	}
	return count, err
}

func (d *Database) CountMessages(ctx context.Context) (int, error) {
	var count int
	err := d.GetContext(ctx, &count, "SELECT COUNT(uid) FROM messages")
	if err != nil {
		return 0, err
	}
	return count, err
}

func (d *Database) CountUnSend(ctx context.Context) (int, error) {
	var count int
	err := d.GetContext(ctx, &count, "SELECT COUNT(uid) FROM messages WHERE status = 0")
	if err != nil {
		return 0, err
	}
	return count, err
}