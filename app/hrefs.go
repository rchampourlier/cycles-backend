// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "cycles": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/rchampourlier/cycles-backend/design
// --out=$(GOPATH)/src/github.com/rchampourlier/cycles-backend
// --version=v1.3.0

package app

// StateHref returns the resource href.
func StateHref() string {
	return "/states/latest"
}
