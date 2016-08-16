package airgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cristianchaparroa/airgo/net"
	"github.com/cristianchaparroa/airgo/response"
)

type AccessToken struct {
	Token string `json:"access_token"`
}

type Airgo struct {
	ApiKey string
}

func NewAPI() *Airgo {
	return &Airgo{}
}

func (a *Airgo) authorize(params *url.Values) (AccessToken, error) {
	var token AccessToken

	resp, err := http.PostForm(Endpoints[Authorize], *params)
	if err != nil {
		return token, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return token, err
	}
	json.Unmarshal(b, &token)
	return token, nil
}

// Login check the user and password in AirBnb
// based on  http://airbnbapi.org/#login-by-email
func (a *Airgo) Login(params *url.Values) (AccessToken, error) {
	params.Add(ClientID, a.ApiKey)
	params.Add(GrantType, GrantTypePassword)
	token, err := a.authorize(params)
	return token, err
}

func (a *Airgo) LoginFB(params *url.Values) (AccessToken, error) {
	params.Add(ClientID, a.ApiKey)
	params.Add(AssertionType, AssertionTypeFacebook)
	params.Add(PreventAccountCreation, PreventAccountCreationTrue)
	token, err := a.authorize(params)
	return token, err
}

func (a *Airgo) LoginGoogle(params *url.Values) (AccessToken, error) {
	params.Add(ClientID, a.ApiKey)
	params.Add(AssertionType, AssertionTypeGoogle)
	params.Add(PreventAccountCreation, PreventAccountCreationTrue)
	token, err := a.authorize(params)
	return token, err
}

// ListingSearch
func (a *Airgo) ListingSearch(params *url.Values) (response.ListingSearchResponse, error) {
	var response response.ListingSearchResponse
	baseUrl, err := url.Parse(Endpoints[Search])
	if err != nil {
		return response, err
	}

	params.Add(ClientID, a.ApiKey)
	baseUrl.RawQuery = params.Encode()

	client := net.HttpClient{}
	headers := http.Header{}
	b, err := client.Get(baseUrl.String(), headers)
	json.Unmarshal(b, &response)
	return response, nil
}

// GetReviews Returns reviews for a given listing.
func (a *Airgo) GetReviews(params *url.Values) (response.ReviewResponse, error) {
	var reviews response.ReviewResponse
	baseUrl, err := url.Parse(Endpoints[GetReviews])
	if err != nil {
		return reviews, err
	}
	params.Add(ClientID, a.ApiKey)
	baseUrl.RawQuery = params.Encode()

	client := net.HttpClient{}
	headers := http.Header{}
	b, err := client.Get(baseUrl.String(), headers)
	json.Unmarshal(b, &reviews)
	return reviews, nil
}

// ViewUserInfo
func (a *Airgo) ViewUserInfo(userID string, params *url.Values) (response.ViewUserInfoResponse, error) {
	var ui response.ViewUserInfoResponse
	u := fmt.Sprintf("%s/%s?", Endpoints[ViewUserInfo], userID)
	baseUrl, err := url.Parse(u)
	if err != nil {
		return ui, err
	}
	params.Add(ClientID, a.ApiKey)
	baseUrl.RawQuery = params.Encode()

	client := net.HttpClient{}
	headers := http.Header{}
	b, err := client.Get(baseUrl.String(), headers)
	json.Unmarshal(b, &ui)
	return ui, nil
}

// ViewListingInfo
func (a *Airgo) ViewListingInfo(listingID string, params *url.Values) (response.ViewListingInfoResponse, error) {
	var li response.ViewListingInfoResponse
	u := fmt.Sprintf("%s/%s?", Endpoints[ViewListingInfo], listingID)

	baseUrl, err := url.Parse(u)
	if err != nil {
		return li, err
	}

	params.Add(ClientID, a.ApiKey)
	baseUrl.RawQuery = params.Encode()

	client := net.HttpClient{}
	headers := http.Header{}
	b, err := client.Get(baseUrl.String(), headers)
	if err != nil {
		return li, err
	}
	json.Unmarshal(b, &li)
	return li, nil
}

func (a *Airgo) CreateMessageThread(token AccessToken, params *url.Values) (response.CreateThreadResponse, error) {
	var response response.CreateThreadResponse
	headers := http.Header{}
	headers.Set("X-Airbnb-OAuth-Token", token.Token)
	params.Add(ClientID, a.ApiKey)

	client := net.HttpClient{}
	b, err := client.Post(Endpoints[CreateMessageThread], *params, headers)

	if err != nil {
		return response, err
	}
	json.Unmarshal(b, &response)
	return response, nil
}

func (a *Airgo) GetMessages(token AccessToken, params *url.Values) {

}

func (a *Airgo) GetUserInfo(token AccessToken, params *url.Values) (response.UserInfoResponse, error) {
	var response response.UserInfoResponse
	baseUrl, err := url.Parse(Endpoints[GetUserInfo])
	if err != nil {
		return response, err
	}

	headers := http.Header{}
	headers.Set("X-Airbnb-OAuth-Token", token.Token)

	params.Add(ClientID, a.ApiKey)
	baseUrl.RawQuery = params.Encode()

	client := net.HttpClient{}
	b, err := client.Get(baseUrl.String(), headers)
	json.Unmarshal(b, &response)
	return response, nil
}
