package items

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type ItemServiceI interface {
	Search(name string, page, limit int) ([]Item, error)
}

type ItemService struct {
	repo ItemRepositoryI // ← pakai interface
}

func NewItemService(repo ItemRepositoryI) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) Search(name string, page, limit int) ([]Item, error) {
	offset := (page - 1) * limit
	return s.repo.Find(name, limit, offset)
}
