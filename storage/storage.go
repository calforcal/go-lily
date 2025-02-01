package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/calforcal/can-lily-eat-it/config"
	"github.com/calforcal/can-lily-eat-it/models"
	_ "github.com/lib/pq"
)

type DBer interface {
	// Add methods as needed, for example:
	GetOrCreateUser(googleUser *models.GoogleUserInfo) (User, error)
}

type PostgresStorage struct {
	db *sql.DB
}

type User struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	GoogleID  string    `json:"google_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewStorage(db *sql.DB) DBer {
	return &PostgresStorage{db: db}
}

func InitDB() (DBer, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DATABASE_HOST,
		config.DATABASE_PORT,
		config.DATABASE_USERNAME,
		config.DATABASE_PASSWORD,
		config.DATABASE_NAME,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	log.Println("Successfully connected to database")
	return NewStorage(db), nil
}

func CloseDB(s DBer) {
	if s != nil {
		s.(*PostgresStorage).db.Close()
	}
}

func (s *PostgresStorage) GetOrCreateUser(googleUser *models.GoogleUserInfo) (User, error) {
	var user User
	err := s.db.QueryRow("SELECT id, uuid, google_id, name, email, created_at, updated_at FROM users WHERE google_id = $1", googleUser.ID).Scan(&user.ID, &user.UUID, &user.GoogleID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			user.UUID = uuid.New().String()
			user.GoogleID = googleUser.ID
			user.Name = googleUser.Name
			user.Email = googleUser.Email
			user.CreatedAt = time.Now()
			user.UpdatedAt = time.Now()

			_, err = s.db.Exec("INSERT INTO users (uuid, google_id, name, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", user.UUID, user.GoogleID, user.Name, user.Email, user.CreatedAt, user.UpdatedAt)
			if err != nil {
				return User{}, fmt.Errorf("error creating user: %v", err)
			}
		} else {
			return User{}, fmt.Errorf("error getting user: %v", err)
		}
	}
	return user, nil
}
