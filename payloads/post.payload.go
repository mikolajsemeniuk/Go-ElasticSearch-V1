package payloads

import (
	"github.com/google/uuid"
)

type Post struct {
	Id      uuid.UUID
	Title   string
	Done    bool
	Created string
	Updated *string
}
