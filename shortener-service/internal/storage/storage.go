package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Storage struct {
	DB *sqlx.DB
}

func NewStorage(connStr string) (*Storage, error) {
	var db *sqlx.DB
	var err error

	for {
		db, err = sqlx.Connect("postgres", connStr)
		if err == nil {
			log.Println("Successfully connected to database")
			break
		}
		log.Printf("Ошибка подключения к базе данных: %v. Повтор через 3 секунды...", err)
		time.Sleep(3 * time.Second)
	}

	return &Storage{DB: db}, nil
}

func (s *Storage) SaveURL(code, original string) error {
	_, err := s.DB.Exec(`INSERT INTO short_urls (shortcode, original_url) VALUES ($1, $2)`, code, original)
	return err
}

func (s *Storage) GetOriginal(code string) (string, error) {
	var url string
	err := s.DB.Get(&url, `SELECT original_url FROM short_urls WHERE shortcode = $1`, code)
	return url, err
}
