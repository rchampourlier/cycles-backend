// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "cycles": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/rchampourlier/cycles-backend/design
// --out=$(GOPATH)/src/github.com/rchampourlier/cycles-backend
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// StateController is the controller interface for the State actions.
type StateController interface {
	goa.Muxer
	Create(*CreateStateContext) error
	Show(*ShowStateContext) error
}

// MountStateController "mounts" a State resource controller on the given service.
func MountStateController(service *goa.Service, ctrl StateController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/states/", ctrl.MuxHandler("preflight", handleStateOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/states/latest", ctrl.MuxHandler("preflight", handleStateOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateStateContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateStatePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleStateOrigin(h)
	service.Mux.Handle("POST", "/states/", ctrl.MuxHandler("create", h, unmarshalCreateStatePayload))
	service.LogInfo("mount", "ctrl", "State", "action", "Create", "route", "POST /states/")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowStateContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleStateOrigin(h)
	service.Mux.Handle("GET", "/states/latest", ctrl.MuxHandler("show", h, nil))
	service.LogInfo("mount", "ctrl", "State", "action", "Show", "route", "GET /states/latest")
}

// handleStateOrigin applies the CORS response headers corresponding to the origin.
func handleStateOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "http://localhost:8080") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Vary", "Origin")
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateStatePayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateStatePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createStatePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
