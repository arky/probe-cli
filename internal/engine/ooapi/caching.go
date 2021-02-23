// Code generated by go generate; DO NOT EDIT.
// 2021-02-23 13:44:06.373936 +0100 CET m=+1.205985560

package ooapi

//go:generate go run ./internal/generator

import (
	"context"
	"reflect"

	"github.com/ooni/probe-cli/v3/internal/engine/ooapi/apimodel"
)

// CheckInCache implements caching for CheckInAPI.
type CheckInCache struct {
	API      CheckInCaller // mandatory
	GobCodec GobCodec      // optional
	KVStore  KVStore       // mandatory
}

type cacheEntryForCheckIn struct {
	Req  *apimodel.CheckInRequest
	Resp *apimodel.CheckInResponse
}

func (c *CheckInCache) Call(ctx context.Context, req *apimodel.CheckInRequest) (*apimodel.CheckInResponse, error) {
	resp, err := c.API.Call(ctx, req)
	if err != nil {
		if resp, _ := c.readcache(req); resp != nil {
			return resp, nil
		}
		return nil, err
	}
	if err := c.writecache(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CheckInCache) gobCodec() GobCodec {
	if c.GobCodec != nil {
		return c.GobCodec
	}
	return &defaultGobCodec{}
}

func (c *CheckInCache) getcache() ([]cacheEntryForCheckIn, error) {
	data, err := c.KVStore.Get("CheckIn.cache")
	if err != nil {
		return nil, err
	}
	var out []cacheEntryForCheckIn
	if err := c.gobCodec().Decode(data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *CheckInCache) setcache(in []cacheEntryForCheckIn) error {
	data, err := c.gobCodec().Encode(in)
	if err != nil {
		return err
	}
	return c.KVStore.Set("CheckIn.cache", data)
}

func (c *CheckInCache) readcache(req *apimodel.CheckInRequest) (*apimodel.CheckInResponse, error) {
	cache, err := c.getcache()
	if err != nil {
		return nil, err
	}
	for _, cur := range cache {
		if reflect.DeepEqual(req, cur.Req) {
			return cur.Resp, nil
		}
	}
	return nil, errCacheNotFound
}

func (c *CheckInCache) writecache(req *apimodel.CheckInRequest, resp *apimodel.CheckInResponse) error {
	cache, _ := c.getcache()
	out := []cacheEntryForCheckIn{{Req: req, Resp: resp}}
	const toomany = 64
	for idx, cur := range cache {
		if reflect.DeepEqual(req, cur.Req) {
			continue // we already updated the cache
		}
		if idx > toomany {
			break
		}
		out = append(out, cur)
	}
	return c.setcache(out)
}

var _ CheckInCaller = &CheckInCache{}

// MeasurementMetaCache implements caching for MeasurementMetaAPI.
type MeasurementMetaCache struct {
	API      MeasurementMetaCaller // mandatory
	GobCodec GobCodec              // optional
	KVStore  KVStore               // mandatory
}

type cacheEntryForMeasurementMeta struct {
	Req  *apimodel.MeasurementMetaRequest
	Resp *apimodel.MeasurementMetaResponse
}

func (c *MeasurementMetaCache) Call(ctx context.Context, req *apimodel.MeasurementMetaRequest) (*apimodel.MeasurementMetaResponse, error) {
	if resp, _ := c.readcache(req); resp != nil {
		return resp, nil
	}
	resp, err := c.API.Call(ctx, req)
	if err != nil {
		return nil, err
	}
	if err := c.writecache(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *MeasurementMetaCache) gobCodec() GobCodec {
	if c.GobCodec != nil {
		return c.GobCodec
	}
	return &defaultGobCodec{}
}

func (c *MeasurementMetaCache) getcache() ([]cacheEntryForMeasurementMeta, error) {
	data, err := c.KVStore.Get("MeasurementMeta.cache")
	if err != nil {
		return nil, err
	}
	var out []cacheEntryForMeasurementMeta
	if err := c.gobCodec().Decode(data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *MeasurementMetaCache) setcache(in []cacheEntryForMeasurementMeta) error {
	data, err := c.gobCodec().Encode(in)
	if err != nil {
		return err
	}
	return c.KVStore.Set("MeasurementMeta.cache", data)
}

func (c *MeasurementMetaCache) readcache(req *apimodel.MeasurementMetaRequest) (*apimodel.MeasurementMetaResponse, error) {
	cache, err := c.getcache()
	if err != nil {
		return nil, err
	}
	for _, cur := range cache {
		if reflect.DeepEqual(req, cur.Req) {
			return cur.Resp, nil
		}
	}
	return nil, errCacheNotFound
}

func (c *MeasurementMetaCache) writecache(req *apimodel.MeasurementMetaRequest, resp *apimodel.MeasurementMetaResponse) error {
	cache, _ := c.getcache()
	out := []cacheEntryForMeasurementMeta{{Req: req, Resp: resp}}
	const toomany = 64
	for idx, cur := range cache {
		if reflect.DeepEqual(req, cur.Req) {
			continue // we already updated the cache
		}
		if idx > toomany {
			break
		}
		out = append(out, cur)
	}
	return c.setcache(out)
}

var _ MeasurementMetaCaller = &MeasurementMetaCache{}

// TestHelpersCache implements caching for TestHelpersAPI.
type TestHelpersCache struct {
	API      TestHelpersCaller // mandatory
	GobCodec GobCodec          // optional
	KVStore  KVStore           // mandatory
}

type cacheEntryForTestHelpers struct {
	Req  *apimodel.TestHelpersRequest
	Resp apimodel.TestHelpersResponse
}

func (c *TestHelpersCache) Call(ctx context.Context, req *apimodel.TestHelpersRequest) (apimodel.TestHelpersResponse, error) {
	resp, err := c.API.Call(ctx, req)
	if err != nil {
		if resp, _ := c.readcache(req); resp != nil {
			return resp, nil
		}
		return nil, err
	}
	if err := c.writecache(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *TestHelpersCache) gobCodec() GobCodec {
	if c.GobCodec != nil {
		return c.GobCodec
	}
	return &defaultGobCodec{}
}

func (c *TestHelpersCache) getcache() ([]cacheEntryForTestHelpers, error) {
	data, err := c.KVStore.Get("TestHelpers.cache")
	if err != nil {
		return nil, err
	}
	var out []cacheEntryForTestHelpers
	if err := c.gobCodec().Decode(data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TestHelpersCache) setcache(in []cacheEntryForTestHelpers) error {
	data, err := c.gobCodec().Encode(in)
	if err != nil {
		return err
	}
	return c.KVStore.Set("TestHelpers.cache", data)
}

func (c *TestHelpersCache) readcache(req *apimodel.TestHelpersRequest) (apimodel.TestHelpersResponse, error) {
	cache, err := c.getcache()
	if err != nil {
		return nil, err
	}
	for _, cur := range cache {
		if reflect.DeepEqual(req, cur.Req) {
			return cur.Resp, nil
		}
	}
	return nil, errCacheNotFound
}

func (c *TestHelpersCache) writecache(req *apimodel.TestHelpersRequest, resp apimodel.TestHelpersResponse) error {
	cache, _ := c.getcache()
	out := []cacheEntryForTestHelpers{{Req: req, Resp: resp}}
	const toomany = 64
	for idx, cur := range cache {
		if reflect.DeepEqual(req, cur.Req) {
			continue // we already updated the cache
		}
		if idx > toomany {
			break
		}
		out = append(out, cur)
	}
	return c.setcache(out)
}

var _ TestHelpersCaller = &TestHelpersCache{}

// TorTargetsCache implements caching for TorTargetsAPI.
type TorTargetsCache struct {
	API      TorTargetsCaller // mandatory
	GobCodec GobCodec         // optional
	KVStore  KVStore          // mandatory
}

type cacheEntryForTorTargets struct {
	Req  *apimodel.TorTargetsRequest
	Resp apimodel.TorTargetsResponse
}

func (c *TorTargetsCache) Call(ctx context.Context, req *apimodel.TorTargetsRequest) (apimodel.TorTargetsResponse, error) {
	resp, err := c.API.Call(ctx, req)
	if err != nil {
		if resp, _ := c.readcache(req); resp != nil {
			return resp, nil
		}
		return nil, err
	}
	if err := c.writecache(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *TorTargetsCache) gobCodec() GobCodec {
	if c.GobCodec != nil {
		return c.GobCodec
	}
	return &defaultGobCodec{}
}

func (c *TorTargetsCache) getcache() ([]cacheEntryForTorTargets, error) {
	data, err := c.KVStore.Get("TorTargets.cache")
	if err != nil {
		return nil, err
	}
	var out []cacheEntryForTorTargets
	if err := c.gobCodec().Decode(data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TorTargetsCache) setcache(in []cacheEntryForTorTargets) error {
	data, err := c.gobCodec().Encode(in)
	if err != nil {
		return err
	}
	return c.KVStore.Set("TorTargets.cache", data)
}

func (c *TorTargetsCache) readcache(req *apimodel.TorTargetsRequest) (apimodel.TorTargetsResponse, error) {
	cache, err := c.getcache()
	if err != nil {
		return nil, err
	}
	for _, cur := range cache {
		if reflect.DeepEqual(req, cur.Req) {
			return cur.Resp, nil
		}
	}
	return nil, errCacheNotFound
}

func (c *TorTargetsCache) writecache(req *apimodel.TorTargetsRequest, resp apimodel.TorTargetsResponse) error {
	cache, _ := c.getcache()
	out := []cacheEntryForTorTargets{{Req: req, Resp: resp}}
	const toomany = 64
	for idx, cur := range cache {
		if reflect.DeepEqual(req, cur.Req) {
			continue // we already updated the cache
		}
		if idx > toomany {
			break
		}
		out = append(out, cur)
	}
	return c.setcache(out)
}

var _ TorTargetsCaller = &TorTargetsCache{}

// URLsCache implements caching for URLsAPI.
type URLsCache struct {
	API      URLsCaller // mandatory
	GobCodec GobCodec   // optional
	KVStore  KVStore    // mandatory
}

type cacheEntryForURLs struct {
	Req  *apimodel.URLsRequest
	Resp *apimodel.URLsResponse
}

func (c *URLsCache) Call(ctx context.Context, req *apimodel.URLsRequest) (*apimodel.URLsResponse, error) {
	resp, err := c.API.Call(ctx, req)
	if err != nil {
		if resp, _ := c.readcache(req); resp != nil {
			return resp, nil
		}
		return nil, err
	}
	if err := c.writecache(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *URLsCache) gobCodec() GobCodec {
	if c.GobCodec != nil {
		return c.GobCodec
	}
	return &defaultGobCodec{}
}

func (c *URLsCache) getcache() ([]cacheEntryForURLs, error) {
	data, err := c.KVStore.Get("URLs.cache")
	if err != nil {
		return nil, err
	}
	var out []cacheEntryForURLs
	if err := c.gobCodec().Decode(data, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *URLsCache) setcache(in []cacheEntryForURLs) error {
	data, err := c.gobCodec().Encode(in)
	if err != nil {
		return err
	}
	return c.KVStore.Set("URLs.cache", data)
}

func (c *URLsCache) readcache(req *apimodel.URLsRequest) (*apimodel.URLsResponse, error) {
	cache, err := c.getcache()
	if err != nil {
		return nil, err
	}
	for _, cur := range cache {
		if reflect.DeepEqual(req, cur.Req) {
			return cur.Resp, nil
		}
	}
	return nil, errCacheNotFound
}

func (c *URLsCache) writecache(req *apimodel.URLsRequest, resp *apimodel.URLsResponse) error {
	cache, _ := c.getcache()
	out := []cacheEntryForURLs{{Req: req, Resp: resp}}
	const toomany = 64
	for idx, cur := range cache {
		if reflect.DeepEqual(req, cur.Req) {
			continue // we already updated the cache
		}
		if idx > toomany {
			break
		}
		out = append(out, cur)
	}
	return c.setcache(out)
}

var _ URLsCaller = &URLsCache{}
