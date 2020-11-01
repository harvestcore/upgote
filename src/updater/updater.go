package updater

import (
	"github.com/google/uuid"
)

type Updater struct {
	Schema   map[interface{}]interface{}
	Interval uint
	Source   string
	Id       uuid.UUID
}

func (u Updater) SendUpdate() {

}

func (u Updater) HandleEvent() {

}

func (u Updater) FetchData() {

}
