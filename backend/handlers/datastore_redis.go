// +build redis

package handlers

import (
	"github.com/mtlynch/whatgotdone/backend/datastore"
	"github.com/mtlynch/whatgotdone/backend/datastore/redis"
)

// New creates a new redis-based Datastore instance.
func newDatastore(datastoreUrl string) datastore.Datastore {
	return redis.New(datastoreUrl)
}
