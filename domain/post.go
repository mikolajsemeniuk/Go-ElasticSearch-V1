package domain

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id      uuid.UUID
	Title   string
	Done    bool
	Created time.Time
	Updated *time.Time
}
