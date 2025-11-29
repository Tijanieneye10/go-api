package todo

import "errors"

type Service struct {
	todos []string
}

func NewService() *Service {
	return &Service{
		todos: make([]string, 0),
	}
}

func (s *Service) Add(todo string) error {
	for _, t := range s.todos {
		if t == todo {
			return errors.New("todo already exist")
		}
	}
	s.todos = append(s.todos, todo)
	return nil
}

func (s *Service) GetAll() []string {
	return s.todos
}
