package models

type UpdateReview struct {
	Id          string `json:"id"`
	PostId      string `json:"post_id"`
	Rating      int    `json:"rating"`
	Description string `json:"description"`
	CustomerId  string `json:"customer_id"`
}
