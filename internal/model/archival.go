package model

//
// Archival format for individual measurement results
// such as TCP connect, TLS handshake, DNS lookup.
//
// These types end up inside the TestKeys field of an
// OONI measurement (see measurement.go).
//
// See https://github.com/ooni/spec/tree/master/data-formats.
//

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"unicode/utf8"
)

//
// Data format extension specification
//

// ArchivalExtSpec describes a data format extension
type ArchivalExtSpec struct {
	Name string // extension name
	V    int64  // extension version
}

// AddTo adds the current ExtSpec to the specified measurement
func (spec ArchivalExtSpec) AddTo(m *Measurement) {
	if m.Extensions == nil {
		m.Extensions = make(map[string]int64)
	}
	m.Extensions[spec.Name] = spec.V
}

var (
	// ArchivalExtDNS is the version of df-002-dnst.md
	ArchivalExtDNS = ArchivalExtSpec{Name: "dnst", V: 0}

	// ArchivalExtNetevents is the version of df-008-netevents.md
	ArchivalExtNetevents = ArchivalExtSpec{Name: "netevents", V: 0}

	// ArchivalExtHTTP is the version of df-001-httpt.md
	ArchivalExtHTTP = ArchivalExtSpec{Name: "httpt", V: 0}

	// ArchivalExtTCPConnect is the version of df-005-tcpconnect.md
	ArchivalExtTCPConnect = ArchivalExtSpec{Name: "tcpconnect", V: 0}

	// ArchivalExtTLSHandshake is the version of df-006-tlshandshake.md
	ArchivalExtTLSHandshake = ArchivalExtSpec{Name: "tlshandshake", V: 0}

	// ArchivalExtTunnel is the version of df-009-tunnel.md
	ArchivalExtTunnel = ArchivalExtSpec{Name: "tunnel", V: 0}
)

//
// Base types
//

// ArchivalBinaryData is a wrapper for bytes that serializes the enclosed
// data using the specific ooni/spec data format for binary data.
//
// See https://github.com/ooni/spec/blob/master/data-formats/df-001-httpt.md#maybebinarydata.
type ArchivalBinaryData []byte

// archivalBinaryDataRepr is the wire representation of binary data according to
// https://github.com/ooni/spec/blob/master/data-formats/df-001-httpt.md#maybebinarydata.
type archivalBinaryDataRepr struct {
	Data   []byte `json:"data"`
	Format string `json:"format"`
}

var (
	_ json.Marshaler   = ArchivalBinaryData{}
	_ json.Unmarshaler = &ArchivalBinaryData{}
)

// MarshalJSON implements json.Marshaler.
func (value ArchivalBinaryData) MarshalJSON() ([]byte, error) {
	// special case: we need to marshal the empty data as the null value
	if len(value) <= 0 {
		return json.Marshal(nil)
	}

	// construct and serialize the OONI representation
	repr := &archivalBinaryDataRepr{Format: "base64", Data: value}
	return json.Marshal(repr)
}

// ErrInvalidBinaryDataFormat is the format returned when marshaling and
// unmarshaling binary data and the value of "format" is unknown.
var ErrInvalidBinaryDataFormat = errors.New("model: invalid binary data format")

// UnmarshalJSON implements json.Unmarshaler.
func (value *ArchivalBinaryData) UnmarshalJSON(raw []byte) error {
	// handle the case where input is a literal null
	if bytes.Equal(raw, []byte("null")) {
		*value = nil
		return nil
	}

	// attempt to unmarshal into the archival representation
	var repr archivalBinaryDataRepr
	if err := json.Unmarshal(raw, &repr); err != nil {
		return err
	}

	// make sure the data format is "base64"
	if repr.Format != "base64" {
		return fmt.Errorf("%w: '%s'", ErrInvalidBinaryDataFormat, repr.Format)
	}

	// we're good because Go uses base64 for []byte automatically
	*value = repr.Data
	return nil
}

// ArchivalMaybeBinaryString is a possibly-binary string. When the string is valid UTF-8
// we serialize it as itself. Otherwise, we use the binary data format defined by
// https://github.com/ooni/spec/blob/master/data-formats/df-001-httpt.md#maybebinarydata
type ArchivalMaybeBinaryString string

var (
	_ json.Marshaler   = ArchivalMaybeBinaryString("")
	_ json.Unmarshaler = (func() *ArchivalMaybeBinaryString { return nil }())
)

// MarshalJSON implements json.Marshaler.
func (value ArchivalMaybeBinaryString) MarshalJSON() ([]byte, error) {
	// convert value to a string
	str := string(value)

	// TODO(bassosimone): here is where we should scrub the string in the future
	// once we have replaced AchivalMaybeBinaryData with ArchivalMaybeBinaryString

	// if we can serialize as UTF-8 string, do that
	if utf8.ValidString(str) {
		return json.Marshal(str)
	}

	// otherwise fallback to the serialization of ArchivalBinaryData
	return json.Marshal(ArchivalBinaryData(str))
}

// UnmarshalJSON implements json.Unmarshaler.
func (value *ArchivalMaybeBinaryString) UnmarshalJSON(rawData []byte) error {
	// first attempt to decode as a string
	var s string
	if err := json.Unmarshal(rawData, &s); err == nil {
		*value = ArchivalMaybeBinaryString(s)
		return nil
	}

	// then attempt to decode as ArchivalBinaryData
	var d ArchivalBinaryData
	if err := json.Unmarshal(rawData, &d); err != nil {
		return err
	}
	*value = ArchivalMaybeBinaryString(d)
	return nil
}

// ArchivalMaybeBinaryData is a possibly binary string. We use this helper class
// to define a custom JSON encoder that allows us to choose the proper
// representation depending on whether the Value field is valid UTF-8 or not.
//
// See https://github.com/ooni/spec/blob/master/data-formats/df-001-httpt.md#maybebinarydata
//
// Deprecated: do not use this type in new code.
type ArchivalMaybeBinaryData struct {
	Value string
}

// MarshalJSON marshals a string-like to JSON following the OONI spec that
// says that UTF-8 content is represented as string and non-UTF-8 content is
// instead represented using `{"format":"base64","data":"..."}`.
func (hb ArchivalMaybeBinaryData) MarshalJSON() ([]byte, error) {
	// if we can serialize as UTF-8 string, do that
	if utf8.ValidString(hb.Value) {
		return json.Marshal(hb.Value)
	}

	// otherwise fallback to the ooni/spec representation for binary data
	er := make(map[string]string)
	er["format"] = "base64"
	er["data"] = base64.StdEncoding.EncodeToString([]byte(hb.Value))
	return json.Marshal(er)
}

// UnmarshalJSON is the opposite of MarshalJSON.
func (hb *ArchivalMaybeBinaryData) UnmarshalJSON(d []byte) error {
	if err := json.Unmarshal(d, &hb.Value); err == nil {
		return nil
	}
	er := make(map[string]string)
	if err := json.Unmarshal(d, &er); err != nil {
		return err
	}
	if v, ok := er["format"]; !ok || v != "base64" {
		return errors.New("missing or invalid format field")
	}
	if _, ok := er["data"]; !ok {
		return errors.New("missing data field")
	}
	b64, err := base64.StdEncoding.DecodeString(er["data"])
	if err != nil {
		return err
	}
	hb.Value = string(b64)
	return nil
}

//
// DNS lookup
//

// ArchivalDNSLookupResult is the result of a DNS lookup.
//
// See https://github.com/ooni/spec/blob/master/data-formats/df-002-dnst.md.
type ArchivalDNSLookupResult struct {
	Answers          []ArchivalDNSAnswer `json:"answers"`
	Engine           string              `json:"engine"`
	Failure          *string             `json:"failure"`
	GetaddrinfoError int64               `json:"getaddrinfo_error,omitempty"`
	Hostname         string              `json:"hostname"`
	QueryType        string              `json:"query_type"`
	RawResponse      []byte              `json:"raw_response,omitempty"`
	Rcode            int64               `json:"rcode,omitempty"`
	ResolverHostname *string             `json:"resolver_hostname"`
	ResolverPort     *string             `json:"resolver_port"`
	ResolverAddress  string              `json:"resolver_address"`
	T0               float64             `json:"t0,omitempty"`
	T                float64             `json:"t"`
	Tags             []string            `json:"tags"`
	TransactionID    int64               `json:"transaction_id,omitempty"`
}

// ArchivalDNSAnswer is a DNS answer.
type ArchivalDNSAnswer struct {
	ASN        int64   `json:"asn,omitempty"`
	ASOrgName  string  `json:"as_org_name,omitempty"`
	AnswerType string  `json:"answer_type"`
	Hostname   string  `json:"hostname,omitempty"`
	IPv4       string  `json:"ipv4,omitempty"`
	IPv6       string  `json:"ipv6,omitempty"`
	TTL        *uint32 `json:"ttl"`
}

//
// TCP connect
//

// ArchivalTCPConnectResult contains the result of a TCP connect.
//
// See https://github.com/ooni/spec/blob/master/data-formats/df-005-tcpconnect.md.
type ArchivalTCPConnectResult struct {
	IP            string                   `json:"ip"`
	Port          int                      `json:"port"`
	Status        ArchivalTCPConnectStatus `json:"status"`
	T0            float64                  `json:"t0,omitempty"`
	T             float64                  `json:"t"`
	Tags          []string                 `json:"tags"`
	TransactionID int64                    `json:"transaction_id,omitempty"`
}

// ArchivalTCPConnectStatus is the status of ArchivalTCPConnectResult.
type ArchivalTCPConnectStatus struct {
	Blocked *bool   `json:"blocked,omitempty"`
	Failure *string `json:"failure"`
	Success bool    `json:"success"`
}

//
// TLS or QUIC handshake
//

// ArchivalTLSOrQUICHandshakeResult is the result of a TLS or QUIC handshake.
//
// See https://github.com/ooni/spec/blob/master/data-formats/df-006-tlshandshake.md
type ArchivalTLSOrQUICHandshakeResult struct {
	Network            string               `json:"network"`
	Address            string               `json:"address"`
	CipherSuite        string               `json:"cipher_suite"`
	Failure            *string              `json:"failure"`
	SoError            *string              `json:"so_error,omitempty"`
	NegotiatedProtocol string               `json:"negotiated_protocol"`
	NoTLSVerify        bool                 `json:"no_tls_verify"`
	PeerCertificates   []ArchivalBinaryData `json:"peer_certificates"`
	ServerName         string               `json:"server_name"`
	T0                 float64              `json:"t0,omitempty"`
	T                  float64              `json:"t"`
	Tags               []string             `json:"tags"`
	TLSVersion         string               `json:"tls_version"`
	TransactionID      int64                `json:"transaction_id,omitempty"`
}

//
// HTTP
//

// ArchivalHTTPRequestResult is the result of sending an HTTP request.
//
// See https://github.com/ooni/spec/blob/master/data-formats/df-001-httpt.md.
type ArchivalHTTPRequestResult struct {
	Network       string               `json:"network,omitempty"`
	Address       string               `json:"address,omitempty"`
	ALPN          string               `json:"alpn,omitempty"`
	Failure       *string              `json:"failure"`
	Request       ArchivalHTTPRequest  `json:"request"`
	Response      ArchivalHTTPResponse `json:"response"`
	T0            float64              `json:"t0,omitempty"`
	T             float64              `json:"t"`
	Tags          []string             `json:"tags"`
	TransactionID int64                `json:"transaction_id,omitempty"`
}

// ArchivalHTTPRequest contains an HTTP request.
//
// Headers are a map in Web Connectivity data format but
// we have added support for a list since January 2020.
type ArchivalHTTPRequest struct {
	Body            ArchivalMaybeBinaryString            `json:"body"`
	BodyIsTruncated bool                                 `json:"body_is_truncated"`
	HeadersList     []ArchivalHTTPHeader                 `json:"headers_list"`
	Headers         map[string]ArchivalMaybeBinaryString `json:"headers"`
	Method          string                               `json:"method"`
	Tor             ArchivalHTTPTor                      `json:"tor"`
	Transport       string                               `json:"x_transport"`
	URL             string                               `json:"url"`
}

// ArchivalHTTPResponse contains an HTTP response.
//
// Headers are a map in Web Connectivity data format but
// we have added support for a list since January 2020.
type ArchivalHTTPResponse struct {
	Body            ArchivalMaybeBinaryString            `json:"body"`
	BodyIsTruncated bool                                 `json:"body_is_truncated"`
	Code            int64                                `json:"code"`
	HeadersList     []ArchivalHTTPHeader                 `json:"headers_list"`
	Headers         map[string]ArchivalMaybeBinaryString `json:"headers"`

	// The following fields are not serialised but are useful to simplify
	// analysing the measurements in telegram, whatsapp, etc.
	Locations []string `json:"-"`
}

// ArchivalNewHTTPHeadersList constructs a new ArchivalHTTPHeader list given HTTP headers.
func ArchivalNewHTTPHeadersList(source http.Header) (out []ArchivalHTTPHeader) {
	out = []ArchivalHTTPHeader{}

	// obtain the header keys
	keys := []string{}
	for key := range source {
		keys = append(keys, key)
	}

	// ensure the output is consistent, which helps with testing;
	// for an example of why we need to sort headers, see
	// https://github.com/ooni/probe-engine/pull/751/checks?check_run_id=853562310
	sort.Strings(keys)

	// insert into the output list
	for _, key := range keys {
		for _, value := range source[key] {
			out = append(out, ArchivalHTTPHeader{
				Key: key,
				Value: ArchivalMaybeBinaryData{
					Value: value,
				},
			})
		}
	}
	return
}

// ArchivalNewHTTPHeadersMap creates a map representation of HTTP headers
func ArchivalNewHTTPHeadersMap(header http.Header) (out map[string]ArchivalMaybeBinaryString) {
	out = make(map[string]ArchivalMaybeBinaryString)
	for key, values := range header {
		for _, value := range values {
			out[key] = ArchivalMaybeBinaryString(value)
			break // just the first header
		}
	}
	return
}

// ArchivalHTTPHeader is a single HTTP header.
type ArchivalHTTPHeader struct {
	Key   string
	Value ArchivalMaybeBinaryData
}

// MarshalJSON marshals a single HTTP header to a tuple where the first
// element is a string and the second element is maybe-binary data.
func (hh ArchivalHTTPHeader) MarshalJSON() ([]byte, error) {
	if utf8.ValidString(hh.Value.Value) {
		return json.Marshal([]string{hh.Key, hh.Value.Value})
	}
	value := make(map[string]string)
	value["format"] = "base64"
	value["data"] = base64.StdEncoding.EncodeToString([]byte(hh.Value.Value))
	return json.Marshal([]interface{}{hh.Key, value})
}

// UnmarshalJSON is the opposite of MarshalJSON.
func (hh *ArchivalHTTPHeader) UnmarshalJSON(d []byte) error {
	var pair []interface{}
	if err := json.Unmarshal(d, &pair); err != nil {
		return err
	}
	if len(pair) != 2 {
		return errors.New("unexpected pair length")
	}
	key, ok := pair[0].(string)
	if !ok {
		return errors.New("the key is not a string")
	}
	value, ok := pair[1].(string)
	if !ok {
		mapvalue, ok := pair[1].(map[string]interface{})
		if !ok {
			return errors.New("the value is neither a string nor a map[string]interface{}")
		}
		if _, ok := mapvalue["format"]; !ok {
			return errors.New("missing format")
		}
		if v, ok := mapvalue["format"].(string); !ok || v != "base64" {
			return errors.New("invalid format")
		}
		if _, ok := mapvalue["data"]; !ok {
			return errors.New("missing data field")
		}
		v, ok := mapvalue["data"].(string)
		if !ok {
			return errors.New("the data field is not a string")
		}
		b64, err := base64.StdEncoding.DecodeString(v)
		if err != nil {
			return err
		}
		value = string(b64)
	}
	hh.Key, hh.Value = key, ArchivalMaybeBinaryData{Value: value}
	return nil
}

// ArchivalHTTPTor contains Tor information.
type ArchivalHTTPTor struct {
	ExitIP   *string `json:"exit_ip"`
	ExitName *string `json:"exit_name"`
	IsTor    bool    `json:"is_tor"`
}

//
// NetworkEvent
//

// ArchivalNetworkEvent is a network event. It contains all the possible fields
// and most fields are optional. They are only added when it makes sense
// for them to be there _and_ we have data to show.
//
// See https://github.com/ooni/spec/blob/master/data-formats/df-008-netevents.md.
type ArchivalNetworkEvent struct {
	Address       string   `json:"address,omitempty"`
	Failure       *string  `json:"failure"`
	NumBytes      int64    `json:"num_bytes,omitempty"`
	Operation     string   `json:"operation"`
	Proto         string   `json:"proto,omitempty"`
	T0            float64  `json:"t0,omitempty"`
	T             float64  `json:"t"`
	TransactionID int64    `json:"transaction_id,omitempty"`
	Tags          []string `json:"tags,omitempty"`
}
