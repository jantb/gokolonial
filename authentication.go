package gokolonial

import (
	"net/url"
	"bytes"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"errors"
)

type User struct {
	IsAuthenticated bool   `json:"is_authenticated"`
	Sessionid       string `json:"sessionid"`
	User struct {
		ID                   int    `json:"id"`
		Email                string `json:"email"`
		FirstName            string `json:"first_name"`
		LastName             string `json:"last_name"`
		MobileNumber         string `json:"mobile_number"`
		BirthDate            string `json:"birth_date"`
		IsProfileComplete    bool   `json:"is_profile_complete"`
		IsDeliveryAvailable  bool   `json:"is_delivery_available"`
		Address1             string `json:"address1"`
		Address2             string `json:"address2"`
		ZipCode              string `json:"zip_code"`
		ZipPlace             string `json:"zip_place"`
		Country              string `json:"country"`
		DeliveryInstructions string `json:"delivery_instructions"`
	} `json:"user"`
}

func (c *Client) Login(username, password string) (user User, err error) {

	obj, err := json.Marshal(Login{Username: username, Password: password})
	if err != nil {
		return user, err
	}
	rel := &url.URL{Path: "/api/v1/user/login/"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(obj))
	if err != nil {
		return user, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("X-Client-Token", c.Token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return user, err
	}
	defer resp.Body.Close()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}

	if resp.StatusCode != 200 {
		return user, errors.New(string(all))
	}
	err = json.Unmarshal(all, &user)
	if err != nil {
		return user, err
	}
	c.cookie = "sessionid=" + user.Sessionid
	return user, nil
}
