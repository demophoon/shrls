package boltstate

import (
	"github.com/demophoon/shrls/pkg/config"

	storm "github.com/asdine/storm/v3"
	log "github.com/sirupsen/logrus"
)

type BoltDBState struct {
	db *storm.DB
}

func New(c *config.Config) *BoltDBState {
	db, err := storm.Open(c.BoltPath)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	state := &BoltDBState{}
	state.db = db
	return state
}
