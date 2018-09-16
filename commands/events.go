package commands

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/rchampourlier/golib"
	"github.com/rchampourlier/golib/s3"
)

// EventsCommands represent the context used to perform
// the queries (e.g. the AWS bucket to be used).
type EventsCommands struct {
	Bucket string
}

// NewEventsCommands returns a valid `EventsCommands` struct.
func NewEventsCommands(bucket string) *EventsCommands {
	return &EventsCommands{Bucket: bucket}
}

// PushEvent pushes a new event that is stored as received
// on S3 and processed to update the event.
func (c *EventsCommands) PushEvent(t time.Time, evt interface{}) (string, error) {
	s3c := s3.NewS3(c.Bucket)

	evtJSON, err := json.Marshal(evt)
	if err != nil {
		return "", err
	}

	key := fmt.Sprintf("events/%s.json", golib.TimestampWithDelimiter(t, "/"))
	err = s3c.CreateObject(key, evtJSON)
	return key, err
}
