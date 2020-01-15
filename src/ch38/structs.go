package ch38

type Request struct {
	TransactionID string `json:"transaction_id"`
	PayLoad []int `json:"payload"`
}

type Response struct {
	TransactionID string `json:"transaction_id"`
	Expression string `json:"exp"`
}

//cmdline
//easyjson -all structs.go