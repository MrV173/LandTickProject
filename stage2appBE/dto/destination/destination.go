package destinationdto

type CreateDestinationRequest struct {
	Name string `json:"name" form:"name"`
}

type DestinationResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
