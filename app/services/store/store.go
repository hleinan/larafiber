package store

import (
	"github.com/gofiber/session/v2"
	"github.com/gofiber/session/v2/provider/sqlite3"
	"log"
	"time"
)

var sessions *session.Session

func init() {
	provider, err := sqlite3.New(sqlite3.Config{
		DBPath:    "database/session.sqlite",
		TableName: "session",
	})

	if err != nil {
		log.Fatal(err.Error())
	}
	sessions = session.New(session.Config{
		Expiration: 24 * time.Hour * 28,
		Provider:   provider,
	})
}

func GetSessions() *session.Session {
	return sessions
}
