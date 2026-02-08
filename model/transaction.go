package model

type Transaction struct {
	ID          int                 `json:"id"`
	TotalAmount int                 `json:"total_amount"`
	CreatedAt   string              `json:"created_at"`
	Details     []TransactionDetail `json:"details,omitempty"`
}

type TransactionDetail struct {
	ID            int    `json:"id"`
	TransactionID int    `json:"transaction_id"`
	ProductID     int    `json:"product_id"`
	ProductName   string `json:"product_name,omitempty"`
	Quantity      int    `json:"quantity"`
	Subtotal      int    `json:"subtotal"`
}

type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}

type SummaryResponse struct {
	TotalRevenue     int             `json:"total_revenue"`
	TotalTransaction int             `json:"total_transaksi"`
	ProductTerlaris  ProductTerlaris `json:"produk_terlaris"`
}

type ProductTerlaris struct {
	Name       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}
