package models

type Order struct {
    ID           int    `json:"id"`
    CustomerName string `json:"customerName"`
    Item         string `json:"item"`
    Quantity     int    `json:"quantity"`
    Status       string `json:"status"`
}
