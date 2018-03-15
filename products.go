package gokolonial

import (
	"net/url"
	"io/ioutil"
	"net/http"
	"errors"
	"encoding/json"
	"fmt"
)

type ProductCategories struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results []struct {
		ID          int         `json:"id"`
		Name        string      `json:"name"`
		Parent      interface{} `json:"parent"`
		Ordering    int         `json:"ordering"`
		Description string      `json:"description"`
		Children []struct {
			ID          int           `json:"id"`
			Name        string        `json:"name"`
			Parent      int           `json:"parent"`
			Ordering    int           `json:"ordering"`
			Description string        `json:"description"`
			Children    []interface{} `json:"children"`
		} `json:"children"`
	} `json:"results"`
}

type ProductsInCategory struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Parent      int           `json:"parent"`
	Ordering    int           `json:"ordering"`
	Description string        `json:"description"`
	Children    []interface{} `json:"children"`
	Products    []struct {
		ID                            int         `json:"id"`
		FullName                      string      `json:"full_name"`
		Brand                         interface{} `json:"brand"`
		Name                          string      `json:"name"`
		NameExtra                     string      `json:"name_extra"`
		FrontURL                      string      `json:"front_url"`
		Images                        interface{} `json:"images"`
		GrossPrice                    string      `json:"gross_price"`
		GrossUnitPrice                string      `json:"gross_unit_price"`
		UnitPriceQuantityAbbreviation string      `json:"unit_price_quantity_abbreviation"`
		UnitPriceQuantityName         string      `json:"unit_price_quantity_name"`
		Discount                      interface{} `json:"discount"`
		Availability                  struct {
			IsAvailable      bool   `json:"is_available"`
			Code             string `json:"code"`
			Description      string `json:"description"`
			DescriptionShort string `json:"description_short"`
		} `json:"availability"`
		CategoryItems []struct {
			ID                int `json:"id"`
			ProductCategoryID int `json:"product_category_id"`
			Ordering          int `json:"ordering"`
		} `json:"category_items"`
		ClientClassifiers []interface{} `json:"client_classifiers"`
	} `json:"products"`
}

type Product struct {
	ID                            int         `json:"id"`
	FullName                      string      `json:"full_name"`
	Brand                         interface{} `json:"brand"`
	Name                          string      `json:"name"`
	NameExtra                     string      `json:"name_extra"`
	FrontURL                      string      `json:"front_url"`
	Images                        interface{} `json:"images"`
	GrossPrice                    string      `json:"gross_price"`
	GrossUnitPrice                string      `json:"gross_unit_price"`
	UnitPriceQuantityAbbreviation string      `json:"unit_price_quantity_abbreviation"`
	UnitPriceQuantityName         string      `json:"unit_price_quantity_name"`
	Discount                      interface{} `json:"discount"`
	Availability                  struct {
		IsAvailable      bool   `json:"is_available"`
		Code             string `json:"code"`
		Description      string `json:"description"`
		DescriptionShort string `json:"description_short"`
	} `json:"availability"`
	CategoryItems []struct {
		ID                int `json:"id"`
		ProductCategoryID int `json:"product_category_id"`
		Ordering          int `json:"ordering"`
	} `json:"category_items"`
	ClientClassifiers []interface{} `json:"client_classifiers"`
	Categories        []struct {
		ID          int           `json:"id"`
		Name        string        `json:"name"`
		Parent      int           `json:"parent"`
		Ordering    int           `json:"ordering"`
		Description string        `json:"description"`
		Children    []interface{} `json:"children"`
	} `json:"categories"`
	ContentsHTML        string `json:"contents_html"`
	NutritionHTML       string `json:"nutrition_html"`
	AlternativeProducts []struct {
		ID                            int         `json:"id"`
		FullName                      string      `json:"full_name"`
		Brand                         interface{} `json:"brand"`
		Name                          string      `json:"name"`
		NameExtra                     string      `json:"name_extra"`
		FrontURL                      string      `json:"front_url"`
		Images                        interface{} `json:"images"`
		GrossPrice                    string      `json:"gross_price"`
		GrossUnitPrice                string      `json:"gross_unit_price"`
		UnitPriceQuantityAbbreviation string      `json:"unit_price_quantity_abbreviation"`
		UnitPriceQuantityName         string      `json:"unit_price_quantity_name"`
		Discount                      interface{} `json:"discount"`
		Availability                  struct {
			IsAvailable      bool   `json:"is_available"`
			Code             string `json:"code"`
			Description      string `json:"description"`
			DescriptionShort string `json:"description_short"`
		} `json:"availability"`
		CategoryItems []struct {
			ID                int `json:"id"`
			ProductCategoryID int `json:"product_category_id"`
			Ordering          int `json:"ordering"`
		} `json:"category_items"`
		ClientClassifiers []interface{} `json:"client_classifiers"`
	} `json:"alternative_products"`
}

func (c *Client) GetProductcategories() (productCategories ProductCategories, err error) {

	rel := &url.URL{Path: "/api/v1/productcategories/"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return productCategories, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("X-Client-Token", c.Token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return productCategories, err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return productCategories, err
	}

	if resp.StatusCode != 200 {
		return productCategories, errors.New(string(all))
	}
	err = json.Unmarshal(all, &productCategories)
	if err != nil {
		return productCategories, err
	}

	return productCategories, nil
}

func (c *Client) GetAllProductsInCategory(id int) (productsInCategory ProductsInCategory, err error) {

	rel := &url.URL{Path: fmt.Sprintf("/api/v1/productcategories/%d/", id)}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return productsInCategory, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("X-Client-Token", c.Token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return productsInCategory, err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return productsInCategory, err
	}

	if resp.StatusCode != 200 {
		return productsInCategory, errors.New(string(all))
	}
	err = json.Unmarshal(all, &productsInCategory)
	if err != nil {
		return productsInCategory, err
	}

	return productsInCategory, nil
}


func (c *Client) GetProduct(id int) (product Product, err error) {

	rel := &url.URL{Path: fmt.Sprintf("/api/v1/products/%d/", id)}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return product, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("X-Client-Token", c.Token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return product, err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return product, err
	}

	if resp.StatusCode != 200 {
		return product, errors.New(string(all))
	}
	err = json.Unmarshal(all, &product)
	if err != nil {
		return product, err
	}

	return product, nil
}
