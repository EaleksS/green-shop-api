package cart

import "github.com/EaleksS/green-shop-api/types"

func findTotalPrice(products []types.Product) float64 {
	var totalPrice float64 = 0

	for _, product := range products {
			totalPrice += product.Price
	}

	return totalPrice
}
