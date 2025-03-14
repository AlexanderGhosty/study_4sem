package main

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	AuthorID   int    `json:"author_id"`
	CategoryID int    `json:"category_id"`
}
