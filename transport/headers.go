package transport

// Edge request context propagation related headers, as defined in
// https://pages.github.snooguts.net/reddit/baseplate.spec/component-apis/thrift#edge-request-context-propagation
const (
	HeaderEdgeRequest = "Edge-Request"
)

// Edge request context propagation related headers, as defined in
// https://pages.github.snooguts.net/reddit/baseplate.spec/component-apis/thrift#edge-request-context-propagation
const (
	// The Trace ID, a 64-bit integer encoded in decimal.
	HeaderTracingTrace = "Trace"
	// The Span ID, a 64-bit integer encoded in decimal.
	HeaderTracingSpan = "Span"
	// The Parent Span ID, a 64-bit integer encoded in decimal.
	HeaderTracingParent = "Parent"
	// The Sampled flag, an ASCII "1" (HeaderTracingSampledTrue) if true,
	// otherwise false.
	// If not present, defaults to false.
	HeaderTracingSampled = "Sampled"
	// Trace flags, a 64-bit integer encoded in decimal.
	// If not present, defaults to null.
	HeaderTracingFlags = "Flags"
)

// HeaderTracingSampledTrue is the header value to indicate that this trace
// should be sampled.
const HeaderTracingSampledTrue = "1"
