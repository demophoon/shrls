package boltstate

import (
	"github.com/demophoon/shrls/pkg/config"

	storm "github.com/asdine/storm/v3"
	log "github.com/sirupsen/logrus"
)

type BoltDBState struct {
	db *storm.DB
}

func New(c config.Config) *BoltDBState {
	// Need to check this isn't nil before continuing
	if c.StateBackend.Bolt == nil {
		log.Fatal("Couldn't initialize BoltDBState backend. Path not defined.")
	}

	log.Tracef("opening up boltdb: %s", c.StateBackend.Bolt.Path)
	db, err := storm.Open(c.StateBackend.Bolt.Path)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	log.Tracef("opened boltdb: %s", c.StateBackend.Bolt.Path)
	state := &BoltDBState{}
	state.db = db
	return state
}

func (s BoltDBState) Close() {
	s.db.Close()
}
