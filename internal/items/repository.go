package items

import "database/sql"

type ItemRepositoryI interface {
	Find(name string, limit, offset int) ([]Item, error)
}

type ItemRepository struct {
	DB *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{DB: db}
}

func (r *ItemRepository) Find(name string, limit, offset int) ([]Item, error) {
	items := []Item{}

	query := `
		SELECT id, name, stock 
		FROM items
		WHERE LOWER(name) LIKE LOWER($1)
		ORDER BY id
		LIMIT $2 OFFSET $3
	`

	rows, err := r.DB.Query(query, "%"+name+"%", limit, offset)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var it Item
		if err := rows.Scan(&it.ID, &it.Name, &it.Stock); err != nil {
			return items, err
		}
		items = append(items, it)
	}

	return items, nil
}
