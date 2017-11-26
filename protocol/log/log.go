package Log

import (
	Network "app/chrome/protocol/network"
	Runtime "app/chrome/protocol/runtime"
)

/*
LogEntry is a log entry.
*/
type LogEntry struct {
	// Log entry source. Allowed values: xml, javascript, network, storage, appcache, rendering,
	// security, deprecation, worker, violation, intervention, recommendation, other.
	Source string `json:"source"`

	// Log entry severity. Allowed values: verbose, info, warning, error.
	Level string `json:"level"`

	// Logged text.
	Text string `json:"text"`

	// Timestamp when this entry was added.
	Timestamp Runtime.Timestamp `json:"timestamp"`

	// Optional. URL of the resource if known.
	URL string `json:"url,omitempty"`

	// Optional. Line number in the resource.
	LineNumber int `json:"lineNumber,omitempty"`

	// Optional. JavaScript stack trace.
	StackTrace *Runtime.StackTrace `json:"stackTrace,omitempty"`

	// Optional. Identifier of the network request associated with this entry.
	NetworkRequestID Network.RequestID `json:"networkRequestId,omitempty"`

	// Optional. Identifier of the worker associated with this entry.
	WorkerID string `json:"workerId,omitempty"`

	// Optional. Call arguments.
	Args []*Runtime.RemoteObject `json:"args,omitempty"`
}

/*
ViolationSetting is a violation configuration setting.
*/
type ViolationSetting struct {
	// Violation type. Allowed values: longTask, longLayout, blockedEvent, blockedParser,
	// discouragedAPIUse, handler, recurringHandler.
	Name string `json:"name"`

	// Time threshold to trigger upon.
	Threshold int `json:"threshold"`
}
