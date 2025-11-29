package todo

type Service struct {
	todos []string
}

func NewService() *Service {
	return &Service{
		todos: make([]string, 0),
	}
}

func (s *Service) Add(todo string) {
	s.todos = append(s.todos, todo)
}

func (s *Service) GetAll() []string {
	return s.todos
}
