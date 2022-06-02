package netx

import (
	"crypto/tls"

	"github.com/ooni/probe-cli/v3/internal/model"
	"github.com/ooni/probe-cli/v3/internal/netxlite"
)

// httpTransportConfig contains the configuration required for constructing an HTTP transport
type httpTransportConfig struct {
	Dialer     model.Dialer
	QUICDialer model.QUICDialer
	TLSDialer  model.TLSDialer
	TLSConfig  *tls.Config
}

// newHTTP3Transport creates a new HTTP3Transport instance.
//
// Deprecation warning
//
// New code should use netxlite.NewHTTP3Transport instead.
func newHTTP3Transport(config httpTransportConfig) model.HTTPTransport {
	// Rationale for using NoLogger here: previously this code did
	// not use a logger as well, so it's fine to keep it as is.
	return netxlite.NewHTTP3Transport(model.DiscardLogger,
		config.QUICDialer, config.TLSConfig)
}

// newSystemTransport creates a new "system" HTTP transport. That is a transport
// using the Go standard library with custom dialer and TLS dialer.
//
// Deprecation warning
//
// New code should use netxlite.NewHTTPTransport instead.
func newSystemTransport(config httpTransportConfig) model.HTTPTransport {
	return netxlite.NewOOHTTPBaseTransport(config.Dialer, config.TLSDialer)
}