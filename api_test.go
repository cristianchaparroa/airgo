package airgo

import (
	"fmt"
	"net/url"
	"testing"
)

func TestLogin(t *testing.T) {
	api := NewAPI()
	api.ApiKey = "3092nxybyb0otqw18e8nh5nty"
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
		t.Errorf("Expected a token but get empty token, %s: ", token.Token)
	}

}

//TODO: Make the test Login with FB
func TestLoginFB(t *testing.T) {
	api := NewAPI()
	api.ApiKey = "3092nxybyb0otqw18e8nh5nty"
	params := &url.Values{}

	//user access token from facebook
	params.Add("assertion", "USER_ACCESS_TOKEN_PROVIDED_BY_FB")

	token, err := api.LoginFB(params)
	fmt.Println(token)
	if err != nil {
		t.Error(err)
	}

	if token.Token == "" {
		t.Errorf("Expected a token but get empty token, %s: ", token.Token)
	}
}

func TestLoginGoogle(t *testing.T) {
	api := NewAPI()
	api.ApiKey = "3092nxybyb0otqw18e8nh5nty"
	params := &url.Values{}

	//user access token from Google
	params.Add("assertion", "USER_ACCESS_TOKEN_PROVIDED_BY_GOOGLE")

	token, err := api.LoginGoogle(params)
	fmt.Println(token)
	if err != nil {
		t.Error(err)
	}

	if token.Token == "" {
		t.Errorf("Expected a token but get empty token, %s: ", token.Token)
	}
}

func TestListingSearch(t *testing.T) {
	api := NewAPI()
	api.ApiKey = "3092nxybyb0otqw18e8nh5nty"

	// TEST 1, full parameters
	params := &url.Values{}
	params.Add(Locale, "en-US")
	params.Add(Currency, "USD")
	params.Add(Format, FormatMinimalPricing)
	params.Add(Limit, "10")
	params.Add(Offset, "0")
	params.Add(FetchFacets, "true")
	params.Add(Guests, "1")
	params.Add(Ib, "false")
	params.Add(IbAddPhotoFlow, "true")
	params.Add(Location, "Lake Tahoe, CA, US")
	params.Add(MinBathRooms, "0")
	params.Add(MinBedRooms, "0")
	params.Add(MinBeds, "1")
	params.Add(MinNumPicUrls, "10")
	params.Add(PriceMax, "210")
	params.Add(PriceMin, "0")
	params.Add(Sort, "1")
	params.Add(UserLat, "37.3398634")
	params.Add(UserLng, "-122.0455164")

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
	api.ApiKey = "3092nxybyb0otqw18e8nh5nty"

	params := &url.Values{}
	_, err := api.GetReviews(params)

	if err != nil {
		t.Error(err)
	}

}

func TestViewUserInfo(t *testing.T)    {}
func TestViewListingInfo(t *testing.T) {}
