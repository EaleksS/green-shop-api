package category

import "github.com/EaleksS/green-shop-api/types"

func findMaxPrice(products []types.Product) float64 {
	if len(products) == 0 {
		return 0
	}

	maxPrice := products[0].Price

	for _, product := range products {
		if product.Price > maxPrice {
			maxPrice = product.Price
		}
	}

	return maxPrice
}
