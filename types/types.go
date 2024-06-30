package types

import "time"

// Products
type ProductStore interface {
	GetProducts() ([]Product, error)
	GetProductsByIDs(ps []int) ([]Product, error)
	GetByIDProducts(id string) (Product, error)
	GetByNameProducts(name string) (Product, error)
	CreateProducts(Product) error
	UpdateProduct(Product) error
}

type Product struct {
	ID 					int `json:"id"`
	Name 				string `json:"name"`
	Category		string `json:"category"`
	Description string `json:"description"`
	Image 			string `json:"image"`
	Price				float64 `json:"price"`
	Quantity 		int `json:"quantity"`
	CreatedAt 	time.Time `json:"createdAt"`
}

type CreateProductPayload struct {
	Name 				string `json:"name" validate:"required"`
	Category		string `json:"category" validate:"required"`
	Description string `json:"description" validate:"required"`
	Image 			string `json:"image" validate:"required"`
	Price				float64 `json:"price" validate:"required"`
	Quantity 		int `json:"quantity" validate:"required"`
}

// Category

type CategoryStore interface {
	GetCategory() ([]Category, error)
	CreateCategory(Category) error
	UpdateCategory(Category) error
}

type Category struct {
	ID 					int `json:"id"`
	Name 				string `json:"name"`
	CreatedAt 	time.Time `json:"createdAt"`
}

type CreateCategoryPayload struct {
	Name 				string `json:"name" validate:"required"`
}

// User

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID        int `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

// Favorite

type FavoriteStore interface {
	GetFavorite(userId int) ([]Favorite, []int, error)
	AddFavorite(userId int, productId int) error
	DeleteFavorite(userId int, productId int) error
}

type Favorite struct {
	ID 					int `json:"id"`
	UserID 			int `json:"userId"`
	ProductID 	int `json:"productId"`
	CreatedAt 	time.Time `json:"createdAt"`
}

type UpdateFavoritePayload struct {
	ID int `json:"id" validate:"required"`
}

// Cart

type CartStore interface {
	GetCart(userId int) ([]Cart, []int, error)
	AddCart(userId int, productId int) error
	DeleteCart(userId int, productId int) error
}

type Cart struct {
	ID 					int `json:"id"`
	UserID 			int `json:"userId"`
	ProductID 	int `json:"productId"`
	CreatedAt 	time.Time `json:"createdAt"`
}

type UpdateCartPayload struct {
	ID int `json:"id" validate:"required"`
}