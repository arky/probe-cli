package measurexlite

//
// Definition of Trace
//

import (
	"time"

	"github.com/miekg/dns"
	"github.com/ooni/probe-cli/v3/internal/model"
	"github.com/ooni/probe-cli/v3/internal/netxlite"
)

// Trace implements model.Trace.
//
// The zero-value of this struct is invalid. To construct you should either
// fill all the fields marked as MANDATORY or use NewTrace.
//
// Buffered channels
//
// NewTrace uses reasonable buffer sizes for the channels used for collecting
// events. You should drain the channels used by this implementation after
// each operation you perform (i.e., we expect you to peform step-by-step
// measurements). If you want larger (or smaller) buffers, then you should
// construct this data type manually with the desired buffer sizes.
//
// We have convenience methods for extracting events from the buffered
// channels. Otherwise, you could read the channels directly. (In which
// case, remember to issue nonblocking channel reads because channels are
// never closed and they're just written when new events occur.)
type Trace struct {
	// Index is the MANDATORY unique index of this trace within the
	// current measurement. If you don't care about uniquely identifying
	// traces, you can use zero to indicate the "default" trace.
	Index int64

	// NetworkEvent is MANDATORY and buffers network events. If you create
	// this channel manually, ensure it has some buffer.
	NetworkEvent chan *model.ArchivalNetworkEvent

	// NewParallelResolverFn is OPTIONAL and can be used to overide
	// calls to the netxlite.NewParallelResolver factory.
	NewParallelResolverFn func() model.Resolver

	// NewDialerWithoutResolverFn is OPTIONAL and can be used to override
	// calls to the netxlite.NewDialerWithoutResolver factory.
	NewDialerWithoutResolverFn func(dl model.DebugLogger) model.Dialer

	// NewTLSHandshakerStdlibFn is OPTIONAL and can be used to overide
	// calls to the netxlite.NewTLSHandshakerStdlib factory.
	NewTLSHandshakerStdlibFn func(dl model.DebugLogger) model.TLSHandshaker

	// DNSLookup is MANDATORY and buffers DNSLookup results based on the
	// query type. When we create this map using NewTrace, we will create
	// an entry for each dns.Type in DNSQueryTypes. If you create this channel
	// manually, you probably want to to the same (and most likely you also
	// want to create buffered channels). Note that the code will print a
	// warning and otherwise ignore all the query types not included in this map.
	DNSLookup map[uint16]chan *model.ArchivalDNSLookupResult

	// TCPConnect is MANDATORY and buffers TCP connect observations. If you create
	// this channel manually, ensure it has some buffer.
	TCPConnect chan *model.ArchivalTCPConnectResult

	// TLSHandshake is MANDATORY and buffers TLS handshake observations. If you create
	// this channel manually, ensure it has some buffer.
	TLSHandshake chan *model.ArchivalTLSOrQUICHandshakeResult

	// TimeNowFn is OPTIONAL and can be used to override calls to time.Now
	// to produce deterministic timing when testing.
	TimeNowFn func() time.Time

	// ZeroTime is the MANDATORY time when we started the current measurement.
	ZeroTime time.Time
}

const (
	// NetworkEventBufferSize is the buffer size for constructing
	// the Trace's NetworkEvent buffered channel.
	NetworkEventBufferSize = 64

	// DNSLookupBufferSize is the buffer size for constructing
	// the Trace's DNSLookup map of buffered channels.
	DNSLookupBufferSize = 8

	// TCPConnectBufferSize is the buffer size for constructing
	// the Trace's TCPConnect buffered channel.
	TCPConnectBufferSize = 8

	// TLSHandshakeBufferSize is the buffer for construcing
	// the Trace's TLSHandshake buffered channel.
	TLSHandshakeBufferSize = 8
)

// DNSQueryTypes contains the list of DNS query types for which
// NewTrace create entries in Trace.DNSLookup.
var DNSQueryTypes = []uint16{
	dns.TypeANY,
	dns.TypeA,
	dns.TypeAAAA,
	dns.TypeCNAME,
	dns.TypeNS,
}

// newDefaultDNSLookupMap is a convenience factory for creating Trace.DNSLookup
func newDefaultDNSLookupMap() map[uint16]chan *model.ArchivalDNSLookupResult {
	out := make(map[uint16]chan *model.ArchivalDNSLookupResult)
	for _, qtype := range DNSQueryTypes {
		out[qtype] = make(chan *model.ArchivalDNSLookupResult, DNSLookupBufferSize)
	}
	return out
}

// NewTrace creates a new instance of Trace using default settings.
//
// We create buffered channels using as buffer sizes the constants that
// are also defined by this package.
//
// Arguments:
//
// - index is the unique index of this trace within the current measurement (use
// zero if you don't care about giving this trace a unique ID);
//
// - zeroTime is the time when we started the current measurement.
func NewTrace(index int64, zeroTime time.Time) *Trace {
	return &Trace{
		Index: index,
		NetworkEvent: make(
			chan *model.ArchivalNetworkEvent,
			NetworkEventBufferSize,
		),
		NewDialerWithoutResolverFn: nil, // use default
		NewTLSHandshakerStdlibFn:   nil, // use default
		DNSLookup:                  newDefaultDNSLookupMap(),
		TCPConnect: make(
			chan *model.ArchivalTCPConnectResult,
			TCPConnectBufferSize,
		),
		TLSHandshake: make(
			chan *model.ArchivalTLSOrQUICHandshakeResult,
			TLSHandshakeBufferSize,
		),
		TimeNowFn: nil, // use default
		ZeroTime:  zeroTime,
	}
}

// newDialerWithoutResolver indirectly calls netxlite.NewDialerWithoutResolver
// thus allowing us to mock this func for testing.
func (tx *Trace) newDialerWithoutResolver(dl model.DebugLogger) model.Dialer {
	if tx.NewDialerWithoutResolverFn != nil {
		return tx.NewDialerWithoutResolverFn(dl)
	}
	return netxlite.NewDialerWithoutResolver(dl)
}

// newParallelResolver indirectly calls the passed netxlite.NewParallerResolver
// thus allowing us to mock this function for testing
func (tx *Trace) newParallelResolver(newResolver func() model.Resolver) model.Resolver {
	if tx.NewParallelResolverFn != nil {
		return tx.NewParallelResolverFn()
	}
	return newResolver()
}

// newTLSHandshakerStdlib indirectly calls netxlite.NewTLSHandshakerStdlib
// thus allowing us to mock this func for testing.
func (tx *Trace) newTLSHandshakerStdlib(dl model.DebugLogger) model.TLSHandshaker {
	if tx.NewTLSHandshakerStdlibFn != nil {
		return tx.NewTLSHandshakerStdlibFn(dl)
	}
	return netxlite.NewTLSHandshakerStdlib(dl)
}

// TimeNow implements model.Trace.TimeNow.
func (tx *Trace) TimeNow() time.Time {
	if tx.TimeNowFn != nil {
		return tx.TimeNowFn()
	}
	return time.Now()
}

// TimeSince is equivalent to Trace.TimeNow().Sub(t0).
func (tx *Trace) TimeSince(t0 time.Time) time.Duration {
	return tx.TimeNow().Sub(t0)
}

var _ model.Trace = &Trace{}