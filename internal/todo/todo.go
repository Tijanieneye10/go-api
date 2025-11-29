package todo

import (
	"errors"
	"strings"
)

type Service struct {
	todos []Item
}

type Item struct {
	Task   string `json:"task"`
	Status bool   `default:"false"json:"status"`
}

func NewService() *Service {
	return &Service{
		todos: make([]Item, 0),
	}
}

func (s *Service) Add(todo string) error {
	for _, t := range s.todos {
		if t.Task == todo {
			return errors.New("todo already exist")
		}
	}
	s.todos = append(s.todos, Item{Task: todo})
	return nil
}

func (s *Service) GetAll() []Item {
	return s.todos
}

func (s *Service) Search(query string) []Item {
	var result []Item

	for _, t := range s.todos {
		if strings.Contains(strings.ToLower(t.Task), strings.ToLower(query)) {
			result = append(result, t)
		}
	}

	return result
}
