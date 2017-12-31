package commands

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/rchampourlier/golib"
)

// StateCommands stores the application context to be used to perform
// the queries (e.g. the AWS bucket to be used).
type StateCommands struct {
	Bucket string
}

// NewStateCommands returns a valid `StateCommands` struct.
func NewStateCommands(bucket string) *StateCommands {
	return &StateCommands{Bucket: bucket}
}

// StoreState stores a new state on the application's storage
// for the specified time.
//
// ### Return values
//
//   - `string`: the identifier of the stored state. This is the
//      identifier than should be used with `DeleteState`.
//   - `error`
//
func (c *StateCommands) StoreState(t time.Time, state interface{}) (string, error) {
	s3 := golib.NewS3(c.Bucket)

	stateJSON, err := json.Marshal(state)
	if err != nil {
		return "", err
	}

	key := fmt.Sprintf("%s.json", golib.TimestampWithDelimiter(t, "/"))
	err = s3.CreateObject(key, stateJSON)
	return key, err
}

// DeleteState deletes the state specified by its identifier (returned by
// `CreateState`.
func (c *StateCommands) DeleteState(stateID string) error {
	s3 := golib.NewS3(c.Bucket)
	err := s3.DeleteObject(stateID)
	return err
}
