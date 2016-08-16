package response

type Facet struct {
	Count int         `json:"count"`
	Key   interface{} `json:"key"`
	Value string      `json:"value"`
}

type Facets struct {
	Availability         []Facet `json:"availability"`
	Bathrooms            []Facet `json:"bathrooms"`
	Bedrooms             []Facet `json:"bedrooms"`
	Beds                 []Facet `json:"beds"`
	CancelPolicy         []Facet `json:"cancel_policy"`
	DescriptionLanguages []Facet `json:"description_languages"`
	HostingAmenityIds    []Facet `json:"hosting_amenity_ids"`
	Languages            []Facet `json:"languages"`
	NeighborhoodFacet    []Facet `json:"neighborhood_facet"`
	PropertyTypeID       []Facet `json:"property_type_id"`
	RoomType             []Facet `json:"room_type"`
	TopAmenities         []Facet `json:"top_amenities"`
	TopNeighborhoods     []Facet `json:"top_neighborhoods"`
}

type Search struct {
	HasAirbnbPlusAvailability bool   `json:"has_airbnb_plus_availability"`
	IsLastMinuteEligible      bool   `json:"is_last_minute_eligible"`
	LastMinuteShowDistSort    bool   `json:"last_minute_show_dist_sort"`
	MobileSessionID           string `json:"mobile_session_id"`
	NativeCurrency            string `json:"native_currency"`
	PriceRangeMaxNative       int    `json:"price_range_max_native"`
	PriceRangeMinNative       int    `json:"price_range_min_native"`
	PriceType                 string `json:"price_type"`
	ResultsCountString        string `json:"results_count_string"`
	SearchID                  string `json:"search_id"`
}

type PriceHistogram struct {
	AveragePrice int   `json:"average_price"`
	Histogram    []int `json:"histogram"`
}

type RationAverage struct {
	Entire_home_apt float64 `json:"Entire home/apt"`
	Private_room    float64 `json:"Private room"`
	Shared_room     float64 `json:"Shared room"`
}
type AvgPriceByRoomType struct {
	AvgPrice interface{}   `json:"avg_price"`
	Ratio    RationAverage `json:"ratio"`
}

type Geography struct {
	Accuracy    int         `json:"accuracy"`
	City        interface{} `json:"city"`
	Country     string      `json:"country"`
	CountryCode string      `json:"country_code"`
	Lat         float64     `json:"lat"`
	Lng         float64     `json:"lng"`
	Precision   string      `json:"precision"`
	State       interface{} `json:"state"`
	StateShort  interface{} `json:"state_short"`
}

type Pagination struct {
	NextOffset  int `json:"next_offset"`
	ResultCount int `json:"result_count"`
}

type MessageCommitment struct {
	Body              interface{} `json:"body"`
	ContextualMessage interface{} `json:"contextual_message"`
	Headline          interface{} `json:"headline"`
}
type UrgencyCommitment struct {
	ListingsLeftAsPercent          interface{}       `json:"listings_left_as_percent"`
	Message                        MessageCommitment `json:"message"`
	MessageType                    interface{}       `json:"message_type"`
	ShowPercentListingsLeftMessage bool              `json:"show_percent_listings_left_message"`
}

type Listing struct {
	AirbnbPlusEnabled     bool          `json:"airbnb_plus_enabled"`
	Bathrooms             int           `json:"bathrooms"`
	Bedrooms              int           `json:"bedrooms"`
	Beds                  int           `json:"beds"`
	City                  string        `json:"city"`
	CoworkerHosted        bool          `json:"coworker_hosted"`
	Distance              interface{}   `json:"distance"`
	ExtraHostLanguages    []interface{} `json:"extra_host_languages"`
	ID                    int           `json:"id"`
	InstantBookable       bool          `json:"instant_bookable"`
	IsBusinessTravelReady bool          `json:"is_business_travel_ready"`
	IsNewListing          bool          `json:"is_new_listing"`
	Lat                   float64       `json:"lat"`
	ListingTags           []interface{} `json:"listing_tags"`
	Lng                   float64       `json:"lng"`
	Name                  string        `json:"name"`
	Neighborhood          string        `json:"neighborhood"`
	PersonCapacity        int           `json:"person_capacity"`
	PictureCount          int           `json:"picture_count"`
	PictureURL            string        `json:"picture_url"`
	PictureUrls           []string      `json:"picture_urls"`
	PrimaryHost           User          `json:"primary_host"`
	PropertyType          string        `json:"property_type"`
	PropertyTypeID        int           `json:"property_type_id"`
	PublicAddress         string        `json:"public_address"`
	ReviewsCount          int           `json:"reviews_count"`
	RoomType              string        `json:"room_type"`
	RoomTypeCategory      string        `json:"room_type_category"`
	StarRating            int           `json:"star_rating"`
	ThumbnailURL          string        `json:"thumbnail_url"`
	User                  User          `json:"user"`
	UserID                int           `json:"user_id"`
	XlPictureURL          string        `json:"xl_picture_url"`
	XlPictureUrls         []string      `json:"xl_picture_urls"`
}

type GuestDetails struct {
	NumberOfAdults   int `json:"number_of_adults"`
	NumberOfChildren int `json:"number_of_children"`
	NumberOfInfants  int `json:"number_of_infants"`
}

type PricingQuoteRatio struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}
type PricingQuote struct {
	Available      bool              `json:"available"`
	CanInstantBook bool              `json:"can_instant_book"`
	CheckIn        interface{}       `json:"check_in"`
	CheckOut       interface{}       `json:"check_out"`
	GuestDetails   GuestDetails      `json:"guest_details"`
	Guests         int               `json:"guests"`
	Price          interface{}       `json:"price"`
	Rate           PricingQuoteRatio `json:"rate"`
	RateType       string            `json:"rate_type"`
}

type SearchResult struct {
	Listing      Listing      `json:"listing"`
	PricingQuote PricingQuote `json:"pricing_quote"`
	ViewedAt     interface{}  `json:"viewed_at"`
}

type Metadata struct {
	AvgPriceByRoomType AvgPriceByRoomType `json:"avg_price_by_room_type"`
	Facets             Facets             `json:"facets"`
	Geography          Geography          `json:"geography"`
	Guidebook          interface{}        `json:"guidebook"`
	ListingTags        struct{}           `json:"listing_tags"`
	ListingsCount      int                `json:"listings_count"`
	Messages           struct{}           `json:"messages"`
	Pagination         Pagination         `json:"pagination"`
	PriceHistogram     PriceHistogram     `json:"price_histogram"`
	Search             Search             `json:"search"`
	UrgencyCommitment  UrgencyCommitment  `json:"urgency_commitment"`
}
type ListingSearchResponse struct {
	Metadata      Metadata       `json:"metadata"`
	SearchResults []SearchResult `json:"search_results"`
}
