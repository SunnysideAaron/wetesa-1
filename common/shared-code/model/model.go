package model

type MetaDataAPIResponse struct {
	Fields  string `json:"fields,omitempty"`
	Filters string `json:"filters,omitempty"`
	Sort    string `json:"sort,omitempty"`
	Page    int    `json:"page,omitempty"`
	Size    int    `json:"size,omitempty"`
	HasNext bool   `json:"hasNext,omitempty"`
}

type MessageAPIResponse struct {
	Severity string `json:"severity"` // should be INFO, WARN, or ERROR
	Message  string `json:"message"`
}
