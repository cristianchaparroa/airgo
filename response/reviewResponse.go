package response

type ReviewMetadata struct {
	ReviewsCount                 int  `json:"reviews_count"`
	ShouldShowReviewTranslations bool `json:"should_show_review_translations"`
}

type ReviewAuthor struct {
	FirstName     string `json:"first_name"`
	HasProfilePic bool   `json:"has_profile_pic"`
	ID            int    `json:"id"`
	PictureURL    string `json:"picture_url"`
	SmartName     string `json:"smart_name"`
	ThumbnailURL  string `json:"thumbnail_url"`
}

type ReviewListing struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ReviewRecipient struct {
	FirstName     string `json:"first_name"`
	HasProfilePic bool   `json:"has_profile_pic"`
	ID            int    `json:"id"`
	PictureURL    string `json:"picture_url"`
	SmartName     string `json:"smart_name"`
	ThumbnailURL  string `json:"thumbnail_url"`
}

type Review struct {
	Author      ReviewAuthor    `json:"author"`
	AuthorID    int             `json:"author_id"`
	CanBeEdited bool            `json:"can_be_edited"`
	Comments    string          `json:"comments"`
	CreatedAt   string          `json:"created_at"`
	ID          int             `json:"id"`
	Listing     ReviewListing   `json:"listing"`
	ListingID   int             `json:"listing_id"`
	Recipient   ReviewRecipient `json:"recipient"`
	RecipientID int             `json:"recipient_id"`
	Response    string          `json:"response"`
	Role        string          `json:"role"`
}

type ReviewResponse struct {
	MetadataReview ReviewMetadata `json:"metadata"`
	Reviews        []Review       `json:"reviews"`
}
