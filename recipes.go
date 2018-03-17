package gokolonial

import (
	"net/url"
	"io/ioutil"
	"net/http"
	"errors"
	"encoding/json"
)


type RecipeTags struct {
	RecipeTags []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"recipe_tags"`
}

type RecipesForTag struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Recipes []struct {
		ID                    int    `json:"id"`
		Title                 string `json:"title"`
		FrontURL              string `json:"front_url"`
		FeatureImageURL       string `json:"feature_image_url"`
		DifficultyString      string `json:"difficulty_string"`
		CookingDurationString string `json:"cooking_duration_string"`
		ConceptIcons          []struct {
			Name     string `json:"name"`
			ImageURL string `json:"image_url"`
		} `json:"concept_icons"`
		IsLikedByUser interface{} `json:"is_liked_by_user"`
	} `json:"recipes"`
}

type Recipe struct {
	ID                    int    `json:"id"`
	Title                 string `json:"title"`
	FrontURL              string `json:"front_url"`
	FeatureImageURL       string `json:"feature_image_url"`
	DifficultyString      string `json:"difficulty_string"`
	CookingDurationString string `json:"cooking_duration_string"`
	ConceptIcons          []struct {
		Name     string `json:"name"`
		ImageURL string `json:"image_url"`
	} `json:"concept_icons"`
	IsLikedByUser      interface{} `json:"is_liked_by_user"`
	ProviderName       string      `json:"provider_name"`
	Lead               string      `json:"lead"`
	PortionPrice       string      `json:"portion_price"`
	DefaultNumPortions int         `json:"default_num_portions"`
	InstructionsHTML   string      `json:"instructions_html"`
	IngredientListHTML string      `json:"ingredient_list_html"`
	Ingredients        []struct {
		ID      int `json:"id"`
		Product struct {
			ID        int    `json:"id"`
			FullName  string `json:"full_name"`
			Brand     string `json:"brand"`
			Name      string `json:"name"`
			NameExtra string `json:"name_extra"`
			FrontURL  string `json:"front_url"`
			Images    []struct {
				Thumbnail struct {
					URL string `json:"url"`
				} `json:"thumbnail"`
				Large struct {
					URL string `json:"url"`
				} `json:"large"`
			} `json:"images"`
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
			ClientClassifiers []struct {
				Name        string `json:"name"`
				ImageURL    string `json:"image_url"`
				IsImportant bool   `json:"is_important"`
			} `json:"client_classifiers"`
		} `json:"product"`
		PortionQuantity        string `json:"portion_quantity"`
		IsBasic                bool   `json:"is_basic"`
		HasAlternativeProducts bool   `json:"has_alternative_products"`
	} `json:"ingredients"`
}

func (c *Client) GetRecipeTags() (recipeTags RecipeTags, err error) {

	rel := &url.URL{Path: "/api/v1/recipe-tags/" }
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return recipeTags, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("X-Client-Token", c.Token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return recipeTags, err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return recipeTags, err
	}

	if resp.StatusCode != 200 {
		return recipeTags, errors.New(string(all))
	}
	err = json.Unmarshal(all, &recipeTags)
	if err != nil {
		return recipeTags, err
	}

	return recipeTags, nil
}

func (c *Client) GetRecipeTagsId(id string) (recipesForTag RecipesForTag, err error) {

	rel := &url.URL{Path: "/api/v1/recipe-tags/" + id}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return recipesForTag, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("X-Client-Token", c.Token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return recipesForTag, err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return recipesForTag, err
	}

	if resp.StatusCode != 200 {
		return recipesForTag, errors.New(string(all))
	}
	err = json.Unmarshal(all, &recipesForTag)
	if err != nil {
		return recipesForTag, err
	}

	return recipesForTag, nil
}
func (c *Client) GetRecipe(id string) (recipe Recipe, err error) {

	rel := &url.URL{Path: "/api/v1/recipes/" + id + "/"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return recipe, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("X-Client-Token", c.Token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return recipe, err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return recipe, err
	}

	if resp.StatusCode != 200 {
		return recipe, errors.New(string(all))
	}
	err = json.Unmarshal(all, &recipe)
	if err != nil {
		return recipe, err
	}

	return recipe, nil
}
