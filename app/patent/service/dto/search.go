package dto

type SimpleSearchReq struct {
	Query string `json:"Query"`
	DB    string `json:"DB"`
}
