package items

type Service struct {
	Repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		Repository: repository,
	}
}

func (s *Service) ListItems() ([]Item, error) {
	return s.Repository.FindAll()
}

func (s *Service) GetItem(id int) (*Item, error) {
	return s.Repository.FindByID(id)
}

func (s *Service) CreateItem(item Item) (*Item, error) {
	return s.Repository.Create(item)
}

func (s *Service) UpdateItem(id int, item Item) (*Item, error) {
	return s.Repository.Update(id, item)
}

func (s *Service) DeleteItem(id int) error {
	return s.Repository.Delete(id)
}
