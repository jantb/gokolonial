package gokolonial

import (
	"net/url"
	"io/ioutil"
	"net/http"
	"errors"
	"encoding/json"
)

type SearchResult struct {
	Products   []Product     `json:"products"`
	Categories []interface{} `json:"categories"`
}

type SearchResultRecipe struct {
	Results []struct {
		ID                    int    `json:"id"`
		Title                 string `json:"title"`
		FrontURL              string `json:"front_url"`
		FeatureImageURL       string `json:"feature_image_url"`
		DifficultyString      string `json:"difficulty_string"`
		CookingDurationString string `json:"cooking_duration_string"`
		ConceptIcons []struct {
			Name     string `json:"name"`
			ImageURL string `json:"image_url"`
		} `json:"concept_icons"`
		IsLikedByUser interface{} `json:"is_liked_by_user"`
	} `json:"results"`
}

func (c *Client) Search(query string) (searchResult SearchResult, err error) {

	rel := &url.URL{Path: "/api/v1/search/?q=" + query}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return searchResult, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("X-Client-Token", c.Token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return searchResult, err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return searchResult, err
	}

	if resp.StatusCode != 200 {
		return searchResult, errors.New(string(all))
	}
	err = json.Unmarshal(all, &searchResult)
	if err != nil {
		return searchResult, err
	}

	return searchResult, nil
}

func (c *Client) SearchRecipes(query string) (searchResultRecipe SearchResultRecipe, err error) {

	rel := &url.URL{Path: "/api/v1/search/recipes/?q=" + query}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return searchResultRecipe, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("X-Client-Token", c.Token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return searchResultRecipe, err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return searchResultRecipe, err
	}

	if resp.StatusCode != 200 {
		return searchResultRecipe, errors.New(string(all))
	}
	err = json.Unmarshal(all, &searchResultRecipe)
	if err != nil {
		return searchResultRecipe, err
	}

	return searchResultRecipe, nil
}
