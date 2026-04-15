package dto

type InvoiceItemRequest struct {
	ItemID   int `json:"item_id"`
	Quantity int `json:"quantity"`
}

type SubmitInvoiceRequest struct {
	SenderName    string                `json:"sender_name"`
	SenderAddress string                `json:"sender_address"`
	ReceiverName  string                `json:"receiver_name"`
	ReceiverAddress string              `json:"receiver_address"`
	Items         []InvoiceItemRequest  `json:"items"`
}