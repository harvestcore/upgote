package db

type FindResponse struct {
	Status bool
	Length int
	Items  []map[string]interface{}
}

type InsertResponse struct {
	Status bool
	Length int
	Items  []interface{}
}

type UpdateResponse struct {
	Status   bool
	Modified int
	Matched  int
}

type DeleteResponse struct {
	Status bool
	Length int
}
