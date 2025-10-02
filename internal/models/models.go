package models

import "github.com/OrbitalJin/michi/internal/sqlc"

type SessionWithUrls struct {
	Session sqlc.Session
	Urls []sqlc.SessionUrl
}
