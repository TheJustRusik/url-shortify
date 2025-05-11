package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Visit struct {
	Shortcode string    `db:"shortcode"`
	IP        string    `db:"ip"`
	URL       string    `db:"url"`
	UserAgent string    `db:"user_agent"`
	VisitedAt time.Time `db:"visited_at"`
}

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

func (s *Storage) SaveVisit(v Visit) error {
	_, err := s.DB.Exec(`
        INSERT INTO visits (shortcode, ip, url, user_agent, visited_at)
        VALUES ($1, $2, $3, $4, $5)
    `, v.Shortcode, v.IP, v.URL, v.UserAgent, v.VisitedAt)
	return err
}

func (s *Storage) GetVisits(shortcode string) ([]Visit, error) {
	var visits []Visit
	err := s.DB.Select(&visits, `
        SELECT shortcode, ip, url, user_agent, visited_at
        FROM visits
        WHERE shortcode = $1
        ORDER BY visited_at DESC
    `, shortcode)
	return visits, err
}
