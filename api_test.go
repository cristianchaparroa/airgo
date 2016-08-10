package airgo

import (
	"fmt"
	"github.com/cristianchaparroa/airgo/config"
	"net/url"
	"testing"
)

func TestLogin(t *testing.T) {
	api := NewAPI()
	c := config.AppConfig()
	api.Setup(c)
	params := &url.Values{}
	params.Add("username", "YOUR_EMAIL")
	params.Add("password", "YOUR_PASSWORD")
	params.Add("locale", "en-US")
	params.Add("currency", "US")

	token, err := api.Login(params)
	fmt.Println(token)
	if err != nil {
		t.Error(err)
	}

	if token.Token == "" {
		t.Errorf("Expected a token but get, %s: ", token.Token)
	}

}

//TODO: Make the test Login with FB
func TestLoginFB(t *testing.T) {}

//TODO: Make the test Login with  Google
func TestLoginGM(t *testing.T) {}

func TestListingSearch(t *testing.T) {
	api := NewAPI()
	c := config.AppConfig()
	api.Setup(c)

	// TEST 1, full parameters
	params := &url.Values{}
	params.Add("locale", "en-US")
	params.Add("currency", "USD")
	params.Add("_format", "for_search_results_with_minimal_pricing")
	params.Add("_limit", "10")
	params.Add("_offset", "0")
	params.Add("fetch_facets", "true")
	params.Add("guests", "1")
	params.Add("ib", "false")
	params.Add("ib_add_photo_flow", "true")
	params.Add("location", "Lake Tahoe, CA, US")
	params.Add("min_bathrooms", "0")
	params.Add("min_bedrooms", "0")
	params.Add("min_beds", "1")
	params.Add("min_num_pic_urls", "10")
	params.Add("price_max", "210")
	params.Add("price_min", "0")
	params.Add("sort", "1")
	params.Add("user_lat", "37.3398634")
	params.Add("user_lng", "-122.0455164")

	// Test 2, Without parameters
	// Test 3, Empty parameters
	//params := &url.Values{}
	//params.Add("locale", "")
	//params.Add("currency", "USD")
	//params.Add("_format", "")
	//params.Add("_limit", "0")
	//params.Add("_offset", "0")
	//params.Add("fetch_facets", "true")
	//params.Add("guests", "0")
	//params.Add("ib", "false")
	//params.Add("ib_add_photo_flow", "false")
	//params.Add("location", "")
	//params.Add("min_bathrooms", "0")
	//params.Add("min_bedrooms", "0")
	//params.Add("min_beds", "0")
	//params.Add("min_num_pic_urls", "10")
	//params.Add("price_max", "210")
	//params.Add("price_min", "0")
	//params.Add("sort", "0")
	//params.Add("user_lat", "")
	//params.Add("user_lng", "")
	_, err := api.ListingSearch(params)

	if err != nil {
		t.Error(err)
	}
}

func TestGetReviews(t *testing.T) {
	api := NewAPI()
	c := config.AppConfig()
	api.Setup(c)

	params := &url.Values{}
	_, err := api.GetReviews(params)

	if err != nil {
		t.Error(err)
	}

}

func TestViewUserInfo(t *testing.T)    {}
func TestViewListingInfo(t *testing.T) {}
