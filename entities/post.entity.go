package entities

import (
	"time"
)

type Post struct {
	Title   string
	Done    bool
	Created time.Time
	Updated *time.Time
}
