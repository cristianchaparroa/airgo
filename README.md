# AIR-GO

This is an **UNOFFICIAL** Airbnb wrapper for GO Lang

### Features

Air-go covers the following features

###### Login Endpoints

- [x] Login with user and pass
- [ ] Login with Facebook
- [ ] Login with Google

###### Get Info Endpoints
- [x] Listing search
- [ ] Get Reviews
- [ ] View User Info
- [ ] View Listing info

###### Get User Endpoints
- [ ] Create message thread
- [ ] Get Messages
- [ ] Get User info

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
c := config.AppConfig()
api.Setup(c)
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

results := api.ListingSearch(params)
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
