package product

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/EaleksS/green-shop-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Category,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)

	if(err != nil) {
		return nil, err
	}

	return product, nil
}

func scanRowIntoProduct(row *sql.Row) (*types.Product, error) {
	product := new(types.Product)

	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Category,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)

	if(err != nil) {
		return nil, err
	}

	return product, nil
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) GetByIDProducts(id string) (types.Product, error) {
	row := s.db.QueryRow("SELECT * FROM products WHERE id IN ($1)", id)
	p, err := scanRowIntoProduct(row)

	if err != nil {
		return types.Product{}, err
	}

	return *p, err
}


func (s *Store) CreateProducts(product types.Product) error {
	_, err := s.db.Exec("INSERT INTO products (name, category, description, image, price, quantity) VALUES ($1, $2, $3, $4, $5, $6)", product.Name, product.Category, product.Description, product.Image, product.Price, product.Quantity)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetProductsByIDs(productIDs []int) ([]types.Product, error) {
	if len(productIDs) == 0 {
		return nil, nil
	}

	placeholders := make([]string, len(productIDs))
	args := make([]interface{}, len(productIDs))
	for i, v := range productIDs {
		args[i] = v
		placeholders[i] = "$" + strconv.Itoa(i+1)
	}

	query := fmt.Sprintf("SELECT  *  FROM products WHERE id IN (%s)", strings.Join(placeholders, ","))
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) GetByNameProducts(name string) (types.Product, error) {
	row := s.db.QueryRow("SELECT * FROM products WHERE name IN ($1)", name)
	p, err := scanRowIntoProduct(row)

	if err != nil {
		return types.Product{}, err
	}

	return *p, err
}


func (s *Store) UpdateProduct(product types.Product) error {
	_, err := s.db.Exec("UPDATE products SET name = ?, category = ?, price = ?, image = ?, description = ?, quantity = ? WHERE id = ?", product.Name, product.Category, product.Price, product.Image, product.Description, product.Quantity, product.ID)
	if err != nil {
		return err
	}

	return nil
}
