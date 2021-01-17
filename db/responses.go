package db

// FindResponse Encapsulates the response of a Find operation
type FindResponse struct {
	Status bool                     `json:"status"`
	Length int                      `json:"length"`
	Items  []map[string]interface{} `json:"items"`
}

// InsertResponse Encapsulates the response of an Insert operation
type InsertResponse struct {
	Status bool          `json:"status"`
	Length int           `json:"length"`
	Items  []interface{} `json:"items"`
}

// UpdateResponse Encapsulates the response of an Update operation
type UpdateResponse struct {
	Status   bool `json:"status"`
	Modified int  `json:"modified"`
	Matched  int  `json:"matched"`
}

// DeleteResponse Encapsulates the response of a Delete operation
type DeleteResponse struct {
	Status bool `json:"status"`
	Length int  `json:"length"`
}
