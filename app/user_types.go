// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "cycles": Application User Types
//
// Command:
// $ goagen
// --design=github.com/rchampourlier/cycles-backend/design
// --out=$(GOPATH)/src/github.com/rchampourlier/cycles-backend
// --version=v1.3.0

package app

// statePayload user type.
type statePayload struct {
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
}

// Publicize creates StatePayload from statePayload
func (ut *statePayload) Publicize() *StatePayload {
	var pub StatePayload
	if ut.State != nil {
		pub.State = ut.State
	}
	return &pub
}

// StatePayload user type.
type StatePayload struct {
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
}
