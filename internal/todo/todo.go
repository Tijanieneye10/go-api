package todo

import (
	"strings"

	"github.com/Tijanieneye10/go-api/internal/database"
)

type Service struct {
	db *database.DB
}

type Item struct {
	Task   string `json:"task"`
	Status bool   `default:"false" json:"status"`
}

func NewService(db *database.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Add(todo string) error {
	_, err := s.db.Sqlite.Exec("INSERT INTO todos (task, status) VALUES (?, ?)", todo, false)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() ([]Item, error) {
	rows, err := s.db.Sqlite.Query("SELECT task, status FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.Task, &item.Status); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (s *Service) Search(query string) ([]Item, error) {
	rows, err := s.db.Sqlite.Query("SELECT task, status FROM todos WHERE LOWER(task) LIKE ?", "%"+strings.ToLower(query)+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.Task, &item.Status); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}
