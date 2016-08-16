package response

type UserInfo struct {
	User User `json:"user"`
}

type ListingUserInfo struct {
	Listing struct {
		City                  string      `json:"city"`
		CollectionIds         interface{} `json:"collection_ids"`
		Country               string      `json:"country"`
		HasDoubleBlindReviews bool        `json:"has_double_blind_reviews"`
		ID                    int         `json:"id"`
		InstantBookable       bool        `json:"instant_bookable"`
		Lat                   float64     `json:"lat"`
		Lng                   float64     `json:"lng"`
		MediumURL             string      `json:"medium_url"`
		Name                  string      `json:"name"`
		NativeCurrency        string      `json:"native_currency"`
		PictureURL            string      `json:"picture_url"`
		Price                 int         `json:"price"`
		PriceFormatted        string      `json:"price_formatted"`
		PriceNative           int         `json:"price_native"`
		SmartLocation         string      `json:"smart_location"`
		ThumbnailURL          string      `json:"thumbnail_url"`
		User                  UserInfo    `json:"user"`
		UserID                int         `json:"user_id"`
		XlPictureURL          string      `json:"xl_picture_url"`
	} `json:"listing"`
}

type ViewUserInfoResponse struct {
	Metadata struct{} `json:"metadata"`
	User     struct {
		About                  string        `json:"about"`
		AcceptanceRate         string        `json:"acceptance_rate"`
		CreatedAt              string        `json:"created_at"`
		FirstName              string        `json:"first_name"`
		FriendsCount           int           `json:"friends_count"`
		HasAvailablePayoutInfo bool          `json:"has_available_payout_info"`
		HasProfilePic          bool          `json:"has_profile_pic"`
		ID                     int           `json:"id"`
		IdentityV2Verified     bool          `json:"identity_v2_verified"`
		IdentityVerified       bool          `json:"identity_verified"`
		IsGeneratedUser        bool          `json:"is_generated_user"`
		IsSuperhost            bool          `json:"is_superhost"`
		Languages              []interface{} `json:"languages"`
		ListingsCount          int           `json:"listings_count"`
		Location               string        `json:"location"`
		Neighborhood           string        `json:"neighborhood"`
		PictureLargeURL        string        `json:"picture_large_url"`
		PictureURL             string        `json:"picture_url"`
		RecentRecommendation   interface{}   `json:"recent_recommendation"`
		RecentReview           RecentReview  `json:"recent_review"`
		RecommendationCount    int           `json:"recommendation_count"`
		ResponseRate           string        `json:"response_rate"`
		ResponseTime           string        `json:"response_time"`
		RevieweeCount          int           `json:"reviewee_count"`
		School                 string        `json:"school"`
		SignupMethod           int           `json:"signup_method"`
		SmartName              string        `json:"smart_name"`
		ThumbnailMediumURL     string        `json:"thumbnail_medium_url"`
		ThumbnailURL           string        `json:"thumbnail_url"`
		TotalListingsCount     int           `json:"total_listings_count"`
		VerificationLabels     []string      `json:"verification_labels"`
		Verifications          []string      `json:"verifications"`
		Work                   string        `json:"work"`
	} `json:"user"`
}
