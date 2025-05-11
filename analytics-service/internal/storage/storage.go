package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type Visit struct {
	Shortcode string    `db:"shortcode"`
	IP        string    `db:"ip"`
	UserAgent string    `db:"user_agent"`
	VisitedAt time.Time `db:"visited_at"`
}

type Storage struct {
	DB *sqlx.DB
}

func NewStorage(dsn string) (*Storage, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &Storage{DB: db}, nil
}

func (s *Storage) SaveVisit(v Visit) error {
	_, err := s.DB.Exec(`
        INSERT INTO visits (shortcode, ip, user_agent, visited_at)
        VALUES ($1, $2, $3, $4)
    `, v.Shortcode, v.IP, v.UserAgent, v.VisitedAt)
	return err
}

func (s *Storage) GetVisits(shortcode string) ([]Visit, error) {
	var visits []Visit
	err := s.DB.Select(&visits, `
        SELECT shortcode, ip, user_agent, visited_at
        FROM visits
        WHERE shortcode = $1
        ORDER BY visited_at DESC
    `, shortcode)
	return visits, err
}
