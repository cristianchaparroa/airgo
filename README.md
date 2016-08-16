# AIR-GO

This is an **UNOFFICIAL** Airbnb wrapper for GO Lang

### Features

Air-go covers the following features

###### Login Endpoints

- [x] Login with user and pass
- [x] Login with Facebook
- [x] Login with Google

###### Get Info Endpoints
- [x] Listing search
- [x] Get Reviews
- [x] View User Info
- [X] View Listing info

###### Get User Endpoints
- [x] Create message thread
- [ ] Get Messages
- [x] Get User info

###### Host Listing  Endpoints

###### Host Messages Endpoints
- [ ] Get Host Messages
- [ ] Pre-approve/Decline
- [ ] Reservation Request
- [ ] Get Reservation Requests
- [ ] Respond To message

###### Host Verification  Endpoints
- [ ] Get Phone Number(s)
- [ ] Submit Verification Code
- [ ] Request Verification Text
- [ ] Request Verification Call

###### Host Payment  Endpoints

- [ ] Set ACH info for Payout


Based on
- http://airbnbapi.org



### Documentation



### Usage

```go
api := airgo.NewApi()
api.ApiKey = "API_kEY"
```

###  Login with email

Returns an access_token, given a valid user account email and password

```go
params := &url.Values{}
params.Add("username", "YOUR_EMAIL")
params.Add("password", "YOUR_PASSWORD")
params.Add("locale", "en-US")
params.Add("currency", "US")

token := api.Login(params)
```

The following are the params allowed to login with email

| Param |      Description    |  Required |
|----------|:-------------:|:------:|
| username |  Account's email address | YES  |
| password |    Account's clear-text password (note: endpoint uses HTTPS)   |   YES |
| locale | Desired language |    NO |
| currency | Currency for listings |    NO |

The following are the Default params  

| Param |      Description    |  Required |
|----------|:-------------:|:------:|
| grant_type|  Required for email authentication (as opposed to OAuth)  | YES  |
| client_id |    API Key   |   YES |

###  Login with Facebook
Returns an access_token, given a valid Facebook user OAuth access token. [See the Facebook docs](https://developers.facebook.com/docs/facebook-login/access-tokens)  to learn how to generate an FB access token.

```go
params := &url.Values{}
//user access token from facebook
params.Add("assertion", "USER_ACCESS_TOKEN_PROVIDED_BY_FB")
//get AirBnb token
token, err := api.LoginFB(params)
```
###  Login with Google
Returns an access_token, given a valid Google user OAuth access token. [See the Google docs](https://developers.google.com/identity/protocols/OAuth2) to learn how to generate a Google access token.
```go
params := &url.Values{}
//user access token from Google
params.Add("assertion", "USER_ACCESS_TOKEN_PROVIDED_BY_GOOGLE")
//get AirBnb token
token, err := api.LoginGoogle(params)
```


### Listing search

Returns listings that fit the given search parameters.
```go
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

listingSearchResp := api.ListingSearch(params)
```
The following are the params allowed to listing search

| Param |      Description    |  Required |
|----------|:-------------:|:------:|
| locale |  Desired lagnuage | NO  |
| currency |    Desired currency  |   NO |
| _format| Search with pricing or not.|    NO |
| _limit | Number of listings to show at a time.|    NO |
| _offset |Number of listings to offset in search. |    NO |
| guests| Number of guests. |    NO |
| ib | Setting to true will only show listings that are instant bookable. |    NO |
| ib_add_photo_flow	 | |    NO |
| location | Search by location name -- if unsure of lat/lng, etc. |    NO |
| min_bathrooms| Minimum number of bathrooms. |    NO |
| min_bedrooms | Minimum number of bedrooms. |    NO |
| min_beds| Minimum number of beds. |    NO |
| price_min | Minimum price. |    NO |
| price_max | Maximum price.|    NO |
| min_num_pic_urls | Minimum number of pictures. |    NO |
| sort | Sorting order (1: forward order, 0: reverse order). |    NO |
| suppress_facets |  |    NO |
| user_lat | Latitude search coordinate.|    NO |
| user_lng | Longitude search coordinate. |    NO |

The following are the Default required params  

| Param |      Description    |  Required |
|----------|:-------------:|:------:|
| client_id |    API Key   |   YES |


### Get Reviews

Returns reviews for a given listing.

```go
params := &url.Values{}
// Required field
params.Add("role", "all")
// required field
params.Add("listing_id", "2056659")

reviewResponse, err := api.GetReviews(params)
```

| Param |      Description    |  Field |
|----------|:-------------:|:------:|
| client_id |  API Key  | Default required |
| role |     |   Default required |
| locale	 |    Desired language |   Optional |
| currency	 |    Desired currency |   optional |
| _format|  |    Optional |
| _limit | Number of reviews to show at a time|    Optional |
| _offset | Number of reviews to offset. |    Optional |


### View User info

Returns detailed information about a user, given his/her/its ID (e.g., found in the view listing endpoint response).

```go
params := &url.Values{}
//if you don't specified the _format you will get a few info
params.Add("_format", "v1_legacy_show")
userId := "37950344"
viewUserInfoResponse, err := api.ViewUserInfo(userId, params)
```


| Param |      Description    |  Field |
|----------|:-------------:|:------:|
| client_id |  API Key  | Default required |
| _format| API result format (just put this -- it'll work without it, but it won't have as much data) |    Optional |
| locale	 |    Desired language |   Optional |
| currency	 |    Desired currency |   optional |



### View Listing info

Returns detailed information about a listing, given its ID (e.g., found in the search endpoint response).

```go
params := &url.Values{}
params.Add("_format", "v1_legacy_for_p3")
listingId := "5116458"
viewListingInfoResponse, err := api.ViewListingInfo(listingId, params)
```

| Param |      Description    |  Field |
|----------|:-------------:|:------:|
| client_id |  API Key  | Default required |
| _format| v1_legacy_for_p3, API result format (just put this -- it won't work without it) |    Default required|
| locale	 |    Desired language |   Optional |
| _source	 |     |   optional |
| number_of_guests		 |  Determines listing availability dates based on the # of guests |   optional |  



### Create Message Thread

Creates a message thread and a stay request, given a valid access token and a listing ID.

NOTE: This is a logged-in endpoint and requires an access_token. See Login Endpoints.


```go

token, err := api.Login(LoginParams)
if err != nil {
  // do something
}

params = &url.Values{}
params.Add("locale", "en-US")
params.Add("currency", "USD")
params.Add("message", "Hi!")
params.Add("checkout_date", "2018-04-02T22:00:00.000-0700")
params.Add("checkin_date", "2018-04-01T00:00:00.000-0700")
params.Add("number_of_guests", "1")
params.Add("listing_id", "10166581")

createdMessageThreadResponse, err := api.CreateMessageThread(token, params)


```

The following are the Form parameters that should be sent.


| Param |      Description    |  Field |
|----------|:-------------:|:------:|
| client_id |  API Key  | Default required |
| listing_id | ID of the listing you'd like to message | Required |
| number_of_guests	 |  Number of guests in the request. | Required |
| checkin_date| Requested check-in date.  | Required |
| checkout_date |  Requested check-out date.  | Required |
| message	 |  Initial message to send (empty to send request only).  | Required |
| locale |  Desired language | Optional|
| currency |  Currency for listings  | Optional |

The following are the header parameters that should be sent.

| Header |      Description    |  Field |
|----------|:-------------:|:------:|
| X-Airbnb-OAuth-Token|  Airbnb auth token (from auth-ing with login endpoints)| Required |


### Get Messages

Returns message threads, given an AirBnB access token (from authenticating with login endpoints).

```go
```
The following are the Form parameters that should be sent.

| Param |      Description    |  Field |
|----------|:-------------:|:------:|
| client_id |  API Key  | Default required |
| locale |  Desired language | Optional|
| offset |  Number of message threads to offset in search  | Optional |
| items_per_page |  Number of message threads to display at once  | Optional |
| role|  Type of threads to retrieve. "guest", "host", or don't include this param for both  | Optional |

The following are the header parameters that should be sent.

| Header |      Description    |  Field |
|----------|:-------------:|:------:|
| X-Airbnb-OAuth-Token|  Airbnb auth token (from auth-ing with login endpoints)| Required |


#### Get User Info

Get basic info about the logged-in user, such as name, picture, phone number, verifications, etc.

```go
token, err := api.Login(LoginParams)

if err != nil {
		//do something
}

params = &url.Values{}
userInfoResponse, err := api.GetUserInfo(token, params)
```
Note: To know more about [UserInfoResponse](#userInfoResponse)

The following are the Form parameters that should be sent.

| Param |      Description    |  Field |
|----------|:-------------:|:------:|
| client_id |  API Key  | Default required |
| locale |  Desired language | Optional|
| offset |  Number of message threads to offset in search  | Optional |
| items_per_page |  Number of message threads to display at once  | Optional |
| role|  Type of threads to retrieve. "guest", "host", or don't include this param for both  | Optional |
| ralert_types[]|  Could be 	`reservation_request`	 | Optional |

The following are the header parameters that should be sent.

| Header |      Description    |  Field |
|----------|:-------------:|:------:|
| X-Airbnb-OAuth-Token|  Airbnb auth token (from auth-ing with login endpoints)| Required |


### Responses

#### <a name="accessTokenResponse"></a> AccessToken
#### ListingSearchResponse  
#### ReviewResponse
#### ViewUserInfoResponse
#### ViewListingInfoResponse
#### CreatedMessageThreadResponse
#### <a name="userInfoResponse"></a> UserInfoResponse
