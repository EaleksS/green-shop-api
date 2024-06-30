package cart

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

func scanRowsIntoCart(rows *sql.Rows) (*types.Cart, error) {
	cart := new(types.Cart)

	err := rows.Scan(
		&cart.ID,
		&cart.UserID,
		&cart.ProductID,
		&cart.CreatedAt,
	)

	if(err != nil) {
		return nil, err
	}

	return cart, nil
}

func (s *Store) GetCart(userId int) ([]types.Cart, []int, error) {
	rows, err := s.db.Query("SELECT * FROM cart WHERE userId IN ($1)", userId)
	if err != nil {
		return nil, nil, err
	}

	cart := make([]types.Cart, 0)
	ids := make([]int, 0)
	for rows.Next() {
		f, err := scanRowsIntoCart(rows)
		if err != nil {
			return nil, nil, err
		}

		cart = append(cart, *f)
		ids = append(ids, *&f.ProductID)
	}

	return cart, ids, nil
}

func (s *Store) AddCart(userId int, productId int)  error {
	_, err := s.db.Query("INSERT INTO cart (userId, productId) VALUES ($1, $2)", userId, productId)
	if err != nil {
		return  err
	}

	return nil
}

func (s *Store) DeleteCart(userId int, productId int) error {
	_, err := s.db.Query("DELETE FROM cart WHERE userId = $1 AND productId = $2", userId, productId)
	if err != nil {
		return  err
	}

	return nil
}