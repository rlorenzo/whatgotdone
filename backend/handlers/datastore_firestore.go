package handlers

import (
	"github.com/mtlynch/whatgotdone/backend/datastore"
	"github.com/mtlynch/whatgotdone/backend/datastore/firestore"
)

// New creates a new firestore-based Datastore instance.
func newDatastore() datastore.Datastore {
	return firestore.New()
}
