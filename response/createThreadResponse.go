package response

type OtherUser struct {
	User User `json:"user"`
}

type Thread struct {
	ID               int       `json:"id"`
	OtherUser        OtherUser `json:"other_user"`
	Preview          string    `json:"preview"`
	PreviewUserID    int       `json:"preview_user_id"`
	PreviouslyUnread bool      `json:"previously_unread"`
	Responded        bool      `json:"responded"`
	Status           string    `json:"status"`
	StatusString     string    `json:"status_string"`
	Unread           bool      `json:"unread"`
	UpdatedAt        string    `json:"updated_at"`
}

type CreateThreadResponse struct {
	Result string `json:"result"`
	Thread struct {
		Thread Thread `json:"thread"`
	} `json:"thread"`
}
