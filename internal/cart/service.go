package cart

type CartServiceI interface {
	Add(itemID, amount int) error
	List() ([]CartItem, error)
	Checkout() error
}

type CartService struct {
	repo CartRepositoryI
}

func NewCartService(repo CartRepositoryI) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) Add(itemID, amount int) error {
	return s.repo.Add(itemID, amount)
}

func (s *CartService) List() ([]CartItem, error) {
	return s.repo.List()
}

func (s *CartService) Checkout() error {
	return s.repo.Clear()
}
