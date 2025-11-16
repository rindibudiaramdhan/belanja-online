package cart

import "database/sql"

type CartRepositoryI interface {
	Add(itemID, amount int) error
	List() ([]CartItem, error)
	Clear() error
}

type CartRepository struct {
	DB *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{DB: db}
}

func (r *CartRepository) Add(itemID, amount int) error {
	_, err := r.DB.Exec(
		`INSERT INTO cart (item_id, amount) VALUES ($1, $2)`,
		itemID, amount,
	)
	return err
}

func (r *CartRepository) List() ([]CartItem, error) {
	rows, err := r.DB.Query(`
		SELECT c.id, c.item_id, c.amount, i.name, i.stock
		FROM cart c
		JOIN items i ON i.id = c.item_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []CartItem

	for rows.Next() {
		var ci CartItem
		if err := rows.Scan(&ci.ID, &ci.Item.ID, &ci.Amount, &ci.Item.Name, &ci.Item.Stock); err != nil {
			return result, err
		}
		result = append(result, ci)
	}

	return result, nil
}

func (r *CartRepository) Clear() error {
	_, err := r.DB.Exec(`DELETE FROM cart`)
	return err
}
