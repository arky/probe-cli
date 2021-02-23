// Code generated by go generate; DO NOT EDIT.
// 2021-02-23 13:44:05.720517 +0100 CET m=+0.552585946

package ooapi

//go:generate go run ./internal/generator

import (
	"bytes"
	"context"
	"net/http"
	"net/url"

	"github.com/ooni/probe-cli/v3/internal/engine/ooapi/apimodel"
)

func (api *CheckReportIDAPI) newRequest(ctx context.Context, req *apimodel.CheckReportIDRequest) (*http.Request, error) {
	URL, err := url.Parse(api.baseURL())
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/_/check_report_id"
	q := url.Values{}
	if req.ReportID == "" {
		return nil, newErrEmptyField("ReportID")
	}
	q.Add("report_id", req.ReportID)
	URL.RawQuery = q.Encode()
	return api.requestMaker().NewRequest(ctx, "GET", URL.String(), nil)
}

func (api *CheckInAPI) newRequest(ctx context.Context, req *apimodel.CheckInRequest) (*http.Request, error) {
	URL, err := url.Parse(api.baseURL())
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/check-in"
	body, err := api.jsonCodec().Encode(req)
	if err != nil {
		return nil, err
	}
	out, err := api.requestMaker().NewRequest(ctx, "POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	out.Header.Set("Content-Type", "application/json")
	return out, nil
}

func (api *LoginAPI) newRequest(ctx context.Context, req *apimodel.LoginRequest) (*http.Request, error) {
	URL, err := url.Parse(api.baseURL())
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/login"
	body, err := api.jsonCodec().Encode(req)
	if err != nil {
		return nil, err
	}
	out, err := api.requestMaker().NewRequest(ctx, "POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	out.Header.Set("Content-Type", "application/json")
	return out, nil
}

func (api *MeasurementMetaAPI) newRequest(ctx context.Context, req *apimodel.MeasurementMetaRequest) (*http.Request, error) {
	URL, err := url.Parse(api.baseURL())
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/measurement_meta"
	q := url.Values{}
	if req.ReportID == "" {
		return nil, newErrEmptyField("ReportID")
	}
	q.Add("report_id", req.ReportID)
	if req.Full {
		q.Add("full", "true")
	}
	if req.Input != "" {
		q.Add("input", req.Input)
	}
	URL.RawQuery = q.Encode()
	return api.requestMaker().NewRequest(ctx, "GET", URL.String(), nil)
}

func (api *RegisterAPI) newRequest(ctx context.Context, req *apimodel.RegisterRequest) (*http.Request, error) {
	URL, err := url.Parse(api.baseURL())
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/register"
	body, err := api.jsonCodec().Encode(req)
	if err != nil {
		return nil, err
	}
	out, err := api.requestMaker().NewRequest(ctx, "POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	out.Header.Set("Content-Type", "application/json")
	return out, nil
}

func (api *TestHelpersAPI) newRequest(ctx context.Context, req *apimodel.TestHelpersRequest) (*http.Request, error) {
	URL, err := url.Parse(api.baseURL())
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/test-helpers"
	return api.requestMaker().NewRequest(ctx, "GET", URL.String(), nil)
}

func (api *PsiphonConfigAPI) newRequest(ctx context.Context, req *apimodel.PsiphonConfigRequest) (*http.Request, error) {
	URL, err := url.Parse(api.baseURL())
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/test-list/psiphon-config"
	return api.requestMaker().NewRequest(ctx, "GET", URL.String(), nil)
}

func (api *TorTargetsAPI) newRequest(ctx context.Context, req *apimodel.TorTargetsRequest) (*http.Request, error) {
	URL, err := url.Parse(api.baseURL())
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/test-list/tor-targets"
	return api.requestMaker().NewRequest(ctx, "GET", URL.String(), nil)
}

func (api *URLsAPI) newRequest(ctx context.Context, req *apimodel.URLsRequest) (*http.Request, error) {
	URL, err := url.Parse(api.baseURL())
	if err != nil {
		return nil, err
	}
	URL.Path = "/api/v1/test-list/urls"
	q := url.Values{}
	if req.CategoryCodes != "" {
		q.Add("category_codes", req.CategoryCodes)
	}
	if req.CountryCode != "" {
		q.Add("country_code", req.CountryCode)
	}
	if req.Limit != 0 {
		q.Add("limit", newQueryFieldInt64(req.Limit))
	}
	URL.RawQuery = q.Encode()
	return api.requestMaker().NewRequest(ctx, "GET", URL.String(), nil)
}

func (api *OpenReportAPI) newRequest(ctx context.Context, req *apimodel.OpenReportRequest) (*http.Request, error) {
	URL, err := url.Parse(api.baseURL())
	if err != nil {
		return nil, err
	}
	URL.Path = "/report"
	body, err := api.jsonCodec().Encode(req)
	if err != nil {
		return nil, err
	}
	out, err := api.requestMaker().NewRequest(ctx, "POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	out.Header.Set("Content-Type", "application/json")
	return out, nil
}

func (api *SubmitMeasurementAPI) newRequest(ctx context.Context, req *apimodel.SubmitMeasurementRequest) (*http.Request, error) {
	URL, err := url.Parse(api.baseURL())
	if err != nil {
		return nil, err
	}
	up, err := api.templateExecutor().Execute("/report/{{ .ReportID }}", req)
	if err != nil {
		return nil, err
	}
	URL.Path = up
	body, err := api.jsonCodec().Encode(req)
	if err != nil {
		return nil, err
	}
	out, err := api.requestMaker().NewRequest(ctx, "POST", URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	out.Header.Set("Content-Type", "application/json")
	return out, nil
}
