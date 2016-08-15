package airgo

const (
	Authorize           = "AUTHORIZE"
	Search              = "SEARCH"
	GetReviews          = "GET_REVIEWS"
	ViewUserInfo        = "VIEW_USER_INFO"
	ViewListingInfo     = "VIEW_LISTING_INFO"
	CreateMessageThread = "CREATE_MESSAGE_THREAD"
	GetMessages         = "GET_MESSAGES"
)

var Endpoints = map[string]string{
	Authorize:           "https://api.airbnb.com/v1/authorize",
	Search:              "https://api.airbnb.com/v2/search_results?",
	GetReviews:          "https://api.airbnb.com/v2/reviews?",
	ViewUserInfo:        "https://api.airbnb.com/v2/users",
	ViewListingInfo:     "https://api.airbnb.com/v2/listings",
	CreateMessageThread: "https://api.airbnb.com/v1/threads/create",
	GetMessages:         "https://api.airbnb.com/v1/threads",
}
