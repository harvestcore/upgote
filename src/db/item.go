package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Item Encapsulates a collection
type Item struct {
	CollectionName string
}

// Collection Returns the collection associated to this class
func (item *Item) Collection() *mongo.Collection {
	var collection *mongo.Collection

	if item.CollectionName != "" {
		collection = DB().Collection(item.CollectionName)
	}

	return collection
}

// Find Find the items that fit the criteria
func (item *Item) Find(criteria map[string]interface{}) *FindResponse {
	var results []map[string]interface{}

	cursor, err := item.Collection().Find(Ctx(), criteria)
	if err == nil {
		err = cursor.All(Ctx(), &results)
	}

	return &FindResponse{
		Status: err == nil,
		Length: len(results),
		Items:  results,
	}
}

// InsertOne Inserts one element in the current collection
func (item *Item) InsertOne(element map[string]interface{}) *InsertResponse {
	res, err := item.Collection().InsertOne(Ctx(), element)
	length := 0
	status := false
	items := make([]interface{}, 0)

	if err == nil {
		length = 1
		status = true
	}

	return &InsertResponse{
		Status: status,
		Length: length,
		Items:  append(items, res),
	}
}

// InsertMany Inserts multiple elements in the current collection
func (item *Item) InsertMany(elements []map[string]interface{}) InsertResponse {
	toInsert := make([]interface{}, 0)

	for _, element := range elements {
		toInsert = append(toInsert, element)
	}

	res, err := item.Collection().InsertMany(Ctx(), toInsert)

	return InsertResponse{
		Status: err == nil,
		Length: len(res.InsertedIDs),
		Items:  res.InsertedIDs,
	}
}

// Update Updates the given element with the new data in the current collection
func (item *Item) Update(criteria map[string]interface{}, updated map[string]interface{}) *UpdateResponse {
	updateQuery := map[string]interface{}{
		"$set": updated,
	}

	res, err := item.Collection().UpdateOne(Ctx(), criteria, updateQuery)
	var modified, matched = 0, 0
	status := false

	if err == nil {
		status = true
		modified = int(res.ModifiedCount)
		matched = int(res.MatchedCount)
	}

	return &UpdateResponse{
		Status:   status,
		Matched:  matched,
		Modified: modified,
	}
}

// Delete Deletes the given element from the current collection
func (item *Item) Delete(criteria map[string]interface{}) *DeleteResponse {
	res, err := item.Collection().DeleteMany(Ctx(), criteria)
	deleted := 0
	status := false

	if err == nil {
		status = true
		deleted = int(res.DeletedCount)
	}

	return &DeleteResponse{
		Status: status,
		Length: deleted,
	}
}
