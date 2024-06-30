package favorite

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

func scanRowsIntoFavorite(rows *sql.Rows) (*types.Favorite, error) {
	favorite := new(types.Favorite)

	err := rows.Scan(
		&favorite.ID,
		&favorite.UserID,
		&favorite.ProductID,
		&favorite.CreatedAt,
	)

	if(err != nil) {
		return nil, err
	}

	return favorite, nil
}

func (s *Store) GetFavorite(userId int) ([]types.Favorite, []int, error) {
	rows, err := s.db.Query("SELECT * FROM favorite WHERE userId IN ($1)", userId)
	if err != nil {
		return nil, nil, err
	}

	favorite := make([]types.Favorite, 0)
	ids := make([]int, 0)
	for rows.Next() {
		f, err := scanRowsIntoFavorite(rows)
		if err != nil {
			return nil, nil, err
		}

		favorite = append(favorite, *f)
		ids = append(ids, *&f.ProductID)
	}

	return favorite, ids, nil
}

func (s *Store) AddFavorite(userId int, productId int)  error {
	_, err := s.db.Query("INSERT INTO favorite (userId, productId) VALUES ($1, $2)", userId, productId)
	if err != nil {
		return  err
	}

	return nil
}

func (s *Store) DeleteFavorite(userId int, productId int) error {
	_, err := s.db.Query("DELETE FROM favorite WHERE userId = $1 AND productId = $2", userId, productId)
	if err != nil {
		return  err
	}

	return nil
}