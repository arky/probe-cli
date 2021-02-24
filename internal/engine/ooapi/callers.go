// Code generated by go generate; DO NOT EDIT.
// 2021-02-24 10:30:25.963482401 +0100 CET m=+0.424715121

package ooapi

//go:generate go run ./internal/generator

import (
	"context"

	"github.com/ooni/probe-cli/v3/internal/engine/ooapi/apimodel"
)

// CheckReportIDCaller represents any type exposing a method
// like CheckReportIDAPI.Call.
type CheckReportIDCaller interface {
	Call(ctx context.Context, req *apimodel.CheckReportIDRequest) (*apimodel.CheckReportIDResponse, error)
}

// CheckInCaller represents any type exposing a method
// like CheckInAPI.Call.
type CheckInCaller interface {
	Call(ctx context.Context, req *apimodel.CheckInRequest) (*apimodel.CheckInResponse, error)
}

// LoginCaller represents any type exposing a method
// like LoginAPI.Call.
type LoginCaller interface {
	Call(ctx context.Context, req *apimodel.LoginRequest) (*apimodel.LoginResponse, error)
}

// MeasurementMetaCaller represents any type exposing a method
// like MeasurementMetaAPI.Call.
type MeasurementMetaCaller interface {
	Call(ctx context.Context, req *apimodel.MeasurementMetaRequest) (*apimodel.MeasurementMetaResponse, error)
}

// RegisterCaller represents any type exposing a method
// like RegisterAPI.Call.
type RegisterCaller interface {
	Call(ctx context.Context, req *apimodel.RegisterRequest) (*apimodel.RegisterResponse, error)
}

// TestHelpersCaller represents any type exposing a method
// like TestHelpersAPI.Call.
type TestHelpersCaller interface {
	Call(ctx context.Context, req *apimodel.TestHelpersRequest) (apimodel.TestHelpersResponse, error)
}

// PsiphonConfigCaller represents any type exposing a method
// like PsiphonConfigAPI.Call.
type PsiphonConfigCaller interface {
	Call(ctx context.Context, req *apimodel.PsiphonConfigRequest) (apimodel.PsiphonConfigResponse, error)
}

// TorTargetsCaller represents any type exposing a method
// like TorTargetsAPI.Call.
type TorTargetsCaller interface {
	Call(ctx context.Context, req *apimodel.TorTargetsRequest) (apimodel.TorTargetsResponse, error)
}

// URLsCaller represents any type exposing a method
// like URLsAPI.Call.
type URLsCaller interface {
	Call(ctx context.Context, req *apimodel.URLsRequest) (*apimodel.URLsResponse, error)
}

// OpenReportCaller represents any type exposing a method
// like OpenReportAPI.Call.
type OpenReportCaller interface {
	Call(ctx context.Context, req *apimodel.OpenReportRequest) (*apimodel.OpenReportResponse, error)
}

// SubmitMeasurementCaller represents any type exposing a method
// like SubmitMeasurementAPI.Call.
type SubmitMeasurementCaller interface {
	Call(ctx context.Context, req *apimodel.SubmitMeasurementRequest) (*apimodel.SubmitMeasurementResponse, error)
}
