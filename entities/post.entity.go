package entities

import (
	"time"
)

type Post struct {
	Title   string
	Done    bool
	Created string
	Updated *time.Time
}
