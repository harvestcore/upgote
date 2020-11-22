package db

// FindResponse Encapsulates the response of a Find operation
type FindResponse struct {
	Status bool
	Length int
	Items  []map[string]interface{}
}

// InsertResponse Encapsulates the response of an Insert operation
type InsertResponse struct {
	Status bool
	Length int
	Items  []interface{}
}

// UpdateResponse Encapsulates the response of an Update operation
type UpdateResponse struct {
	Status   bool
	Modified int
	Matched  int
}

// DeleteResponse Encapsulates the response of a Delete operation
type DeleteResponse struct {
	Status bool
	Length int
}
