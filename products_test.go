package gokolonial

import (
	"testing"
)

func TestGetProductcategories(t *testing.T) {
	err, client := NewClient()
	productCategories, err := client.GetProductcategories()
	if err != nil {
		t.Fail()
	}
	if productCategories.Count == 0 {
		t.Fail()
	}
	products, err := client.GetAllProductsInCategory(productCategories.Results[0].ID)
	if len(products.Children) == 0 {
		t.Fail()
	}
}
