// Package firestore implements a datastore.Datastore interface using Google
// Cloud Firestore as a backend.
package firestore

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"

	"github.com/mtlynch/whatgotdone/backend/types"
)

type (
	client struct {
		firestoreClient *firestore.Client
		ctx             context.Context
	}

	userDocument struct {
		Username     string `firestore:"username,omitempty"`
		LastModified string `firestore:"lastModified,omitempty"`
	}

	reactionsDocument struct {
		Reactions []types.Reaction `firestore:"reactions,omitempty"`
	}

	entryReactionsDocument struct {
		entryAuthor string `firestore:"entryAuthor,omitempty"`
		entryDate   string `firestore:"entryDate,omitempty"`
	}
)

const (
	entriesRootKey      = "journalEntries"
	perUserEntriesKey   = "entries"
	draftsRootKey       = "journalDrafts"
	perUserDraftsKey    = "drafts"
	reactionsRootKey    = "reactions"
	perUserReactionsKey = "perUserReactions"
	secretsRootKey      = "secrets"
	secretUserKitDocKey = "userKitKey"
)

func getGoogleCloudProjectId() string {
	projectId := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectId == "" {
		log.Fatalf("GOOGLE_CLOUD_PROJECT environment variable must be set")
	}
	return projectId
}
