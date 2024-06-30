package product

import (
	"sort"
	"strings"

	"github.com/EaleksS/green-shop-api/types"
)

func reverseArray(arr []types.Product) []types.Product {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func sortBySort(ps []types.Product, sortBy string) []types.Product {
	if sortBy == "" {
		return ps
	}

	products := ps

	if sortBy == "ascending" {
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price < products[j].Price
		})
	}

	if sortBy == "descending" {
		sort.Slice(products, func(i, j int) bool {
			return products[i].Price > products[j].Price
		})
	}

	return products
}

func priceSort(ps []types.Product, highPrice float64, lowPrice float64) []types.Product {
	if highPrice == 0 {
		return ps
	}

	products := make([]types.Product, 0)

	for _, p := range ps {
		if p.Price >= lowPrice && p.Price <= highPrice {
			products = append(products, p)
		}
	}
	
	return products
}

func categorySort(ps []types.Product, category string) []types.Product {
	if category == "" {
		return ps
	}

	products := make([]types.Product, 0)

	for _, p := range ps {
		if strings.ToLower(p.Category) == strings.ToLower(category) {
			products = append(products, p)
		}
	}
	
	return products
}

func searchSort(ps []types.Product, search string) []types.Product {
	if search == "" {
		return ps
	}

	products := make([]types.Product, 0)

	for _, p := range ps {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(search)) {
			products = append(products, p)
		}
	}
	
	return products
}