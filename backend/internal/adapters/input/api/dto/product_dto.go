package dto

type ProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type GetProductsResponse struct {
	Data  []ProductResponse `json:"data"`
	Page  int               `json:"_page"`
	Limit int               `json:"_limit"`
	Total int64             `json:"_total"`
	Links PaginationLinks   `json:"_links"`
}

type GetProductByIDResponse struct {
	Data  ProductResponse `json:"data"`
	Links []LinkDTO       `json:"_links"`
}

type LinkDTO struct {
	Rel    string `json:"rel"`
	Href   string `json:"href"`
	Method string `json:"method"`
}

type PaginationLinks struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}
