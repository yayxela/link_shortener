package entities

type AddLinkRequest struct {
	URL string `json:"url" validate:"required,http_url"`
}

type AddLinkResponse struct {
	ShortURL string `json:"url"`
}
