package dtos

type BookDto struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type BookRequest struct {
	Title string `json:"title"`
}
