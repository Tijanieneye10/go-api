package todo

import "errors"

type Service struct {
	todos []Item
}

type Item struct {
	Task   string
	Status bool
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
	s.todos = append(s.todos, Item{Task: todo,
		Status: true,
	})
	return nil
}

func (s *Service) GetAll() []Item {
	return s.todos
}
