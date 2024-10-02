package dto

type ProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type GetProductsResponse struct {
	Products []ProductResponse `json:"products"`
	Page     int               `json:"_page"`
	Limit    int               `json:"_limit"`
	Total    int64             `json:"_total"`
}

type GetProductByIDResponse struct {
	ProductResponse
}

type CreateProductRequest struct {
	ProductResponse
}

type CreateProductResponse struct {
	ProductResponse
}
