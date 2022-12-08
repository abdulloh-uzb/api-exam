package models

type CreatePost struct {
	OwnerId     string  `json:"owner_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Content     string  `json:"content"`
	Medias      []Media `json:"medias"`
}

type Media struct {
	Name string `json:"name"`
	Link string `json:"link"`
	Type string `json:"type"`
}

type UpdatePost struct {
	Id          string  `json:"id"`
	OwnerId     string  `json:"owner_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Content     string  `json:"content"`
	Medias      []UpdateMedia `json:"medias"`
}

type UpdateMedia struct {
	Id     string `json:"id"`
	PostId string `json:"post_id"`
	Name   string `json:"name"`
	Link   string `json:"link"`
	Type   string `json:"type"`
}
