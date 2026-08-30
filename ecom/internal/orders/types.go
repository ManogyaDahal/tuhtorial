package orders

type OrderItem struct { 
	ProductID int64 `json:"productId"`
	Quantity int32 `json:"quantity"`
}

type createOrderParams struct { 
	CustomerId int64 	`json:"customerId"`
	Items []OrderItem `json:"items"`
}
