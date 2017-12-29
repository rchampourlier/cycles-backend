package queries_test

import (
	"os"
	"testing"
	"time"

	"github.com/rchampourlier/cycles-backend/commands"
	"github.com/rchampourlier/cycles-backend/queries"
)

type context struct {
	bucket   string
	commands *commands.StateCommands
	queries  *queries.StateQueries
	t        *testing.T
}

func prepareContext(t *testing.T) *context {
	bucket := os.Getenv("AWS_BUCKET_TEST")
	if len(bucket) == 0 {
		t.Errorf("AWS_BUCKET_TEST must be defined")
	}
	return &context{
		bucket:   bucket,
		commands: commands.NewStateCommands(bucket),
		queries:  queries.NewStateQueries(bucket),
		t:        t,
	}
}

func createState(ctx *context, year int, month time.Month, day int, state string) string {
	t := time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
	stateID, err := ctx.commands.StoreState(t, state)
	if err != nil {
		ctx.t.Errorf(err.Error())
	}
	return stateID
}

func deleteState(ctx *context, stateID string) {
	err := ctx.commands.DeleteState(stateID)
	if err != nil {
		ctx.t.Errorf("could not delete state `%s` (%s)", stateID, err.Error())
	}
}

func getLatestState(ctx *context) string {
	s, err := ctx.queries.GetLatestState()
	if err != nil {
		ctx.t.Errorf(err.Error())
	}
	return s
}

func TestGetLatestState(t *testing.T) {
	ctx := prepareContext(t)
	id1 := createState(ctx, 2015, time.January, 1, "state-1")
	id2 := createState(ctx, 2017, time.January, 1, "state-2")
	id3 := createState(ctx, 2017, time.February, 1, "state-3")
	latestState := getLatestState(ctx)
	for _, item := range []string{id1, id2, id3} {
		deleteState(ctx, item)
	}
	expectedState := "state-3"
	if latestState != expectedState {
		t.Errorf("Expected latest state to be `%s`, got `%s`", expectedState, latestState)
	}
}
