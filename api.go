package airgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cristianchaparroa/airgo/config"
	"github.com/cristianchaparroa/airgo/net"
	"github.com/cristianchaparroa/airgo/response"
)

type AccessToken struct {
	Token string `json:"access_token"`
}

type Airgo struct {
	Conf config.Config
}

func NewAPI() *Airgo {
	return &Airgo{}
}

func (a *Airgo) Setup(c config.Config) {
	a.Conf = c
}

// Login check the user and password in AirBnb
// based on  http://airbnbapi.org/#login-by-email
func (a *Airgo) Login(params *url.Values) (AccessToken, error) {
	var token AccessToken

	params.Add("client_id", a.Conf.ClientID)
	params.Add("grant_type", "password")

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

// TODO: Implements Login with facebook
func (a *Airgo) LoginFB(params *url.Values) (AccessToken, error) {
	var token AccessToken

	params.Add("client_id", a.Conf.ClientID)
	//Required for Facebook authentication.
	params.Add("assertion_type", "https://graph.facebook.com/me")
	//For sign-in, as opposed to registration.
	params.Add("prevent_account_creation", "true")
	
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

// TODO: Implements Login wiht Gmail
func (a *Airgo) LoginGM() {}

// ListingSearch
func (a *Airgo) ListingSearch(params *url.Values) (response.ListingSearchResp, error) {
	var response response.ListingSearchResp
	baseUrl, err := url.Parse(Endpoints[Search])
	if err != nil {
		return response, err
	}

	params.Add("client_id", a.Conf.ClientID)
	baseUrl.RawQuery = params.Encode()

	client := net.HttpClient{}
	b, err := client.Get(baseUrl.String())
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
	params.Add("client_id", a.Conf.ClientID)
	baseUrl.RawQuery = params.Encode()

	client := net.HttpClient{}
	b, err := client.Get(baseUrl.String())
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
	params.Add("client_id", a.Conf.ClientID)
	baseUrl.RawQuery = params.Encode()

	client := net.HttpClient{}
	b, err := client.Get(baseUrl.String())
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

	params.Add("client_id", a.Conf.ClientID)
	baseUrl.RawQuery = params.Encode()

	client := net.HttpClient{}
	b, err := client.Get(baseUrl.String())
	json.Unmarshal(b, &li)
	return li, nil
}
