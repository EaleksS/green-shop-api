package category

import (
	"database/sql"

	"github.com/EaleksS/green-shop-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func scanRowsIntoCategory(rows *sql.Rows) (*types.Category, error) {
	category := new(types.Category)

	err := rows.Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
	)

	if(err != nil) {
		return nil, err
	}

	return category, nil
}

func (s *Store) GetCategory() ([]types.Category, error) {
	rows, err := s.db.Query("SELECT * FROM category")
	if err != nil {
		return nil, err
	}

	category := make([]types.Category, 0)
	for rows.Next() {
		p, err := scanRowsIntoCategory(rows)
		if err != nil {
			return nil, err
		}

		category = append(category, *p)
	}

	return category, nil
}

func (s *Store) CreateCategory(category types.Category) error {
	_, err := s.db.Exec("INSERT INTO category (name) VALUES ($1)", category.Name)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateCategory(category types.Category) error {
	_, err := s.db.Exec("UPDATE category SET name = ? WHERE id = ?", category.Name, category.ID)
	if err != nil {
		return err
	}

	return nil
}
