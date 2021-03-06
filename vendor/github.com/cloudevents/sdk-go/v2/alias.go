package v2

// Package cloudevents alias' common functions and types to improve discoverability and reduce
// the number of imports for simple HTTP clients.

import (
	"github.com/cloudevents/sdk-go/v2/binding"
	"github.com/cloudevents/sdk-go/v2/client"
	"github.com/cloudevents/sdk-go/v2/context"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/observability"
	"github.com/cloudevents/sdk-go/v2/protocol"
	"github.com/cloudevents/sdk-go/v2/protocol/http"
	"github.com/cloudevents/sdk-go/v2/types"
)

// Client

type ClientOption client.Option
type Client = client.Client

// Event

type Event = event.Event

// Context

type EventContext = event.EventContext
type EventContextV1 = event.EventContextV1
type EventContextV03 = event.EventContextV03

// Custom Types

type Timestamp = types.Timestamp
type URIRef = types.URIRef

// HTTP Protocol

type HTTPOption http.Option

type HTTPProtocol = http.Protocol

// Encoding

type Encoding = binding.Encoding

// Message

type Message = binding.Message

const (
	// ReadEncoding

	ApplicationXML                  = event.ApplicationXML
	ApplicationJSON                 = event.ApplicationJSON
	TextPlain                       = event.TextPlain
	ApplicationCloudEventsJSON      = event.ApplicationCloudEventsJSON
	ApplicationCloudEventsBatchJSON = event.ApplicationCloudEventsBatchJSON
	Base64                          = event.Base64

	// Event Versions

	VersionV1  = event.CloudEventsVersionV1
	VersionV03 = event.CloudEventsVersionV03

	// Encoding

	EncodingBinary     = binding.EncodingBinary
	EncodingStructured = binding.EncodingStructured
)

var (

	// ContentType Helpers

	StringOfApplicationJSON                 = event.StringOfApplicationJSON
	StringOfApplicationXML                  = event.StringOfApplicationXML
	StringOfTextPlain                       = event.StringOfTextPlain
	StringOfApplicationCloudEventsJSON      = event.StringOfApplicationCloudEventsJSON
	StringOfApplicationCloudEventsBatchJSON = event.StringOfApplicationCloudEventsBatchJSON
	StringOfBase64                          = event.StringOfBase64

	// Client Creation

	NewClient             = client.New
	NewClientObserved     = client.NewObserved
	NewDefaultClient      = client.NewDefault
	NewHTTPReceiveHandler = client.NewHTTPReceiveHandler

	// Client Options

	WithEventDefaulter   = client.WithEventDefaulter
	WithUUIDs            = client.WithUUIDs
	WithTimeNow          = client.WithTimeNow
	WithTracePropagation = client.WithTracePropagation()

	// Event Creation

	NewEvent  = event.New
	NewResult = protocol.NewResult

	NewHTTPResult = http.NewResult

	// Message Creation

	ToMessage = binding.ToMessage

	// HTTP Messages

	WriteHTTPRequest = http.WriteRequest

	// Tracing

	EnableTracing = observability.EnableTracing

	// Context

	ContextWithTarget      = context.WithTarget
	TargetFromContext      = context.TargetFrom
	WithEncodingBinary     = binding.WithForceBinary
	WithEncodingStructured = binding.WithForceStructured

	// Custom Types

	ParseTimestamp = types.ParseTimestamp
	ParseURIRef    = types.ParseURIRef
	ParseURI       = types.ParseURI

	// HTTP Protocol

	NewHTTP = http.New

	// HTTP Protocol Options

	WithTarget          = http.WithTarget
	WithHeader          = http.WithHeader
	WithShutdownTimeout = http.WithShutdownTimeout
	//WithEncoding           = http.WithEncoding
	//WithStructuredEncoding = http.WithStructuredEncoding // TODO: expose new way
	WithPort         = http.WithPort
	WithPath         = http.WithPath
	WithMiddleware   = http.WithMiddleware
	WithListener     = http.WithListener
	WithRoundTripper = http.WithRoundTripper
)
