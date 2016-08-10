package response

type User struct {
	FirstName     string `json:"first_name"`
	HasProfilePic bool   `json:"has_profile_pic"`
	ID            int    `json:"id"`
	IsSuperhost   bool   `json:"is_superhost"`
	PictureURL    string `json:"picture_url"`
	SmartName     string `json:"smart_name"`
	ThumbnailURL  string `json:"thumbnail_url"`
	RevieweeCount int    `json:"reviewee_count"`
}

type Reviewer struct {
	User User `json:"user"`
}

type ReviewInfo struct {
	Comments  string   `json:"comments"`
	CreatedAt string   `json:"created_at"`
	ID        int      `json:"id"`
	Listing   struct{} `json:"listing"`
	ListingID int      `json:"listing_id"`
	Reviewee  struct {
		User User `json:"user"`
	} `json:"reviewee"`
	RevieweeID int      `json:"reviewee_id"`
	Reviewer   Reviewer `json:"reviewer"`
	ReviewerID int      `json:"reviewer_id"`
	Role       string   `json:"role"`
}

type RecentReview struct {
	Review ReviewInfo `json:"review"`
}
