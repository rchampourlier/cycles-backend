package queries

import (
	"fmt"

	"github.com/rchampourlier/golib"
)

// StateQueries stores the application context to be used to perform
// the queries (e.g. the AWS bucket to be used).
type StateQueries struct {
	Bucket string
}

// NewStateQueries returns a valid `StateQueries` struct.
//
// ### Params
//
//   - `bucket string`: the name of the AWS bucket to be used.
//
func NewStateQueries(bucket string) *StateQueries {
	return &StateQueries{Bucket: bucket}
}

// GetLatestState returns the latest stored state. The timestamp
// used as a key is used to determine which one is the latest.
func (q *StateQueries) GetLatestState() (string, error) {
	s3 := golib.NewS3(q.Bucket)
	key, err := s3.FindLatestInTimestampPrefixedObjects("/")
	if err != nil {
		return "", err
	}
	if key == nil {
		return "", fmt.Errorf("not found")
	}
	contents, err := s3.FetchObject(*key)
	return string(contents), err
}
