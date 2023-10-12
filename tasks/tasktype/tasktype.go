package tasktype

import "time"

const (
	DeletedImage = "deleted:image"
)

type DeletedImagePayload struct {
	Time time.Time
}
