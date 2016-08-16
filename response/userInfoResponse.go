package response

type UserInfoResponse struct {
	AnalyticsProperties struct {
		AddedToWishlist        interface{} `json:"added_to_wishlist"`
		AffiliateID            int         `json:"affiliate_id"`
		BornDevice             interface{} `json:"born_device"`
		BornUser               string      `json:"born_user"`
		City                   interface{} `json:"city"`
		FacebookConnected      bool        `json:"facebook_connected"`
		HaveHosted             interface{} `json:"have_hosted"`
		HaveReviewed           interface{} `json:"have_reviewed"`
		HaveTraveled           interface{} `json:"have_traveled"`
		IsActiveHost           interface{} `json:"is_active_host"`
		Market                 interface{} `json:"market"`
		NumAcceptedResos       interface{} `json:"num_accepted_resos"`
		NumberOfActiveListings interface{} `json:"number_of_active_listings"`
		NumberOfListings       interface{} `json:"number_of_listings"`
		ReceivedInquiry        interface{} `json:"received_inquiry"`
	} `json:"analytics_properties"`
	BadgeCounts struct {
		HostGroups  int `json:"host_groups"`
		HostHome    int `json:"host_home"`
		TravelInbox int `json:"travel_inbox"`
	} `json:"badge_counts"`
	Config struct {
		BraintreeBlue bool `json:"braintree_blue"`
	} `json:"config"`
	Currency                string `json:"currency"`
	IsVrPlatformPoweredHost bool   `json:"is_vr_platform_powered_host"`
	LatestVersionCode       int    `json:"latest_version_code"`
	LatestVersionNumber     string `json:"latest_version_number"`
	Locale                  string `json:"locale"`
	PhotoExperiment         bool   `json:"photo_experiment"`
	Result                  string `json:"result"`
	TosAgreementRequired    bool   `json:"tos_agreement_required"`
	Unread                  struct {
		AlertsCount   int `json:"alerts_count"`
		MessagesCount int `json:"messages_count"`
	} `json:"unread"`
	User struct {
		User struct {
			About                         string        `json:"about"`
			AcceptanceRate                string        `json:"acceptance_rate"`
			Birthdate                     string        `json:"birthdate"`
			CountryOfResidence            string        `json:"country_of_residence"`
			CreatedAt                     string        `json:"created_at"`
			Email                         string        `json:"email"`
			FirstName                     string        `json:"first_name"`
			FriendsCount                  int           `json:"friends_count"`
			Gender                        string        `json:"gender"`
			Groups                        string        `json:"groups"`
			HasAvailablePayoutInfo        bool          `json:"has_available_payout_info"`
			HasProfilePic                 bool          `json:"has_profile_pic"`
			ID                            int           `json:"id"`
			IdentityV2Verified            bool          `json:"identity_v2_verified"`
			IdentityVerified              bool          `json:"identity_verified"`
			IsContentManager              bool          `json:"is_content_manager"`
			IsGeneratedUser               bool          `json:"is_generated_user"`
			IsSuperhost                   bool          `json:"is_superhost"`
			Languages                     []interface{} `json:"languages"`
			LastName                      string        `json:"last_name"`
			ListingsCount                 int           `json:"listings_count"`
			Location                      interface{}   `json:"location"`
			PhoneNumbers                  []interface{} `json:"phone_numbers"`
			PictureURL                    string        `json:"picture_url"`
			PubliclyVisibleWishlistsCount int           `json:"publicly_visible_wishlists_count"`
			RecommendationCount           int           `json:"recommendation_count"`
			ResponseRate                  string        `json:"response_rate"`
			ResponseTime                  string        `json:"response_time"`
			RevieweeCount                 int           `json:"reviewee_count"`
			School                        string        `json:"school"`
			SignupMethod                  int           `json:"signup_method"`
			ThumbnailMediumURL            string        `json:"thumbnail_medium_url"`
			ThumbnailURL                  string        `json:"thumbnail_url"`
			Timezone                      string        `json:"timezone"`
			TotalListingsCount            int           `json:"total_listings_count"`
			VerificationLabels            []interface{} `json:"verification_labels"`
			Verifications                 []interface{} `json:"verifications"`
			WishlistsCount                int           `json:"wishlists_count"`
			Work                          string        `json:"work"`
		} `json:"user"`
	} `json:"user"`
}
