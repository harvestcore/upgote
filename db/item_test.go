package db_test

import (
	"testing"

	"github.com/harvestcore/upgote/db"
	"github.com/harvestcore/upgote/types"
	"github.com/stretchr/testify/assert"
)

var testingCollection = "testing"

func TestItemCreation(t *testing.T) {
	var item = db.Item{
		CollectionName: testingCollection,
	}
	assert.NotNil(t, item, "Item creation return nil with all parameters set")
}

func TestItemInsertOne(t *testing.T) {
	var name = types.Dict{
		"name": "pi",
	}

	var item = db.Item{
		CollectionName: testingCollection,
	}

	res := item.InsertOne(name)
	assert.Equal(t, res.Status, true, "Status insert is not true")
	assert.Equal(t, res.Length, 1, "Single item not inserted")
	assert.Equal(t, len(res.Items), 1, "Only 1 item not inserted")

	found := item.Find(name)
	assert.Equal(t, found.Status, true, "Status find is not true")
	assert.Equal(t, found.Length, 1, "Should only find one item")
	assert.Equal(t, len(found.Items), 1, "Should only find one item")

	item.Drop()
}

func TestItemInsertMany(t *testing.T) {
	names := []types.Dict{
		{
			"test": "TEST",
		},
		{
			"test": "TEST",
		},
	}

	name := types.Dict{
		"test": "TEST",
	}

	var item = db.Item{
		CollectionName: testingCollection,
	}

	res := item.InsertMany(names)
	assert.Equal(t, res.Status, true, "Status insert is not true")
	assert.Equal(t, res.Length, 2, "Single item not inserted")
	assert.Equal(t, len(res.Items), 2, "Only 2 items not inserted")

	found := item.Find(name)
	assert.Equal(t, found.Status, true, "Status find is not true")
	assert.Equal(t, found.Length, 2, "Should only find 2 items")
	assert.Equal(t, len(found.Items), 2, "Should only find 2 items")

	item.Drop()
}

func TestItemFind(t *testing.T) {
	names := []types.Dict{
		{
			"test": "TEST",
		},
		{
			"test": "TEST2",
		},
	}

	name := types.Dict{
		"test": "TEST2",
	}

	empty := types.Dict{}

	var item = db.Item{
		CollectionName: testingCollection,
	}

	res := item.InsertMany(names)
	assert.Equal(t, res.Status, true, "Status insert is not true")
	assert.Equal(t, res.Length, 2, "Single item not inserted")
	assert.Equal(t, len(res.Items), 2, "Only 2 items not inserted")

	var found = item.Find(name)
	assert.Equal(t, found.Status, true, "Status find is not true")
	assert.Equal(t, found.Length, 1, "Should only find one item")
	assert.Equal(t, len(found.Items), 1, "Should only find one item")

	found = item.Find(empty)
	assert.Equal(t, found.Status, true, "Status find is not true")
	assert.Equal(t, found.Length, 2, "Should only find 2 items")
	assert.Equal(t, len(found.Items), 2, "Should only find 2 items")

	item.Drop()
}

func TestItemDelete(t *testing.T) {
	names := []types.Dict{
		{
			"test": "TEST",
		},
		{
			"test": "TEST",
		},
	}

	name := types.Dict{
		"test": "TEST",
	}

	empty := types.Dict{}

	var item = db.Item{
		CollectionName: testingCollection,
	}

	res := item.InsertMany(names)
	assert.Equal(t, res.Status, true, "Status insert is not true")
	assert.Equal(t, res.Length, 2, "Single item not inserted")
	assert.Equal(t, len(res.Items), 2, "Only 2 items not inserted")

	deleted := item.Delete(name)
	assert.Equal(t, deleted.Status, true, "Status delete is not true")
	assert.Equal(t, deleted.Length, 2, "Should delete 2 items")

	found := item.Find(empty)
	assert.Equal(t, found.Status, true, "Status find is not true")
	assert.Equal(t, found.Length, 0, "Should not find items")
	assert.Equal(t, len(found.Items), 0, "Should not find items")

	item.Drop()
}

func TestItemUpdate(t *testing.T) {
	var name = types.Dict{
		"name": "pi",
	}

	nameUpdate := types.Dict{
		"name":    "pi",
		"surname": "po",
	}

	empty := types.Dict{}

	var item = db.Item{
		CollectionName: testingCollection,
	}

	res := item.InsertOne(name)
	assert.Equal(t, res.Status, true, "Status insert is not true")
	assert.Equal(t, res.Length, 1, "Single item not inserted")
	assert.Equal(t, len(res.Items), 1, "Only 1 item not inserted")

	updated := item.Update(name, nameUpdate)
	assert.Equal(t, updated.Status, true, "Status find is not true")
	assert.Equal(t, updated.Matched, 1, "Should only match one item")
	assert.Equal(t, updated.Modified, 1, "Should only modify one item")

	found := item.Find(empty)
	assert.Equal(t, found.Status, true, "Status find is not true")
	assert.Equal(t, found.Length, 1, "Should find 1 item")
	assert.Equal(t, len(found.Items), 1, "Should find 1 item")
	assert.NotNil(t, found.Items[0], "Item should exist")
	assert.Equal(t, found.Items[0]["name"], "pi", "Name should be pi")
	assert.Equal(t, found.Items[0]["surname"], "po", "Surname should be po")

	item.Drop()
}
