// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/bjxujiang/gf.

// Package gtrace provides convenience wrapping functionality for tracing feature using OpenTelemetry.
package gtrace

import (
	"context"
	"github.com/bjxujiang/gf/container/gmap"
	"github.com/bjxujiang/gf/container/gvar"
	"github.com/bjxujiang/gf/net/gipv4"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"
	"os"
	"strings"
)

const (
	tracingCommonKeyIpIntranet = `ip.intranet`
	tracingCommonKeyIpHostname = `hostname`
)

var (
	intranetIps, _ = gipv4.GetIntranetIpArray()
	intranetIpStr  = strings.Join(intranetIps, ",")
	hostname, _    = os.Hostname()
)

// CommonLabels returns common used attribute labels:
// ip.intranet, hostname.
func CommonLabels() []label.KeyValue {
	return []label.KeyValue{
		label.String(tracingCommonKeyIpHostname, hostname),
		label.String(tracingCommonKeyIpIntranet, intranetIpStr),
	}
}

// IsActivated checks and returns if tracing feature is activated.
func IsActivated(ctx context.Context) bool {
	return GetTraceId(ctx) != ""
}

// GetTraceId retrieves and returns TraceId from context.
// It returns an empty string is tracing feature is not activated.
func GetTraceId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	traceId := trace.SpanContextFromContext(ctx).TraceID
	if traceId.IsValid() {
		return traceId.String()
	}
	return ""
}

// GetSpanId retrieves and returns SpanId from context.
// It returns an empty string is tracing feature is not activated.
func GetSpanId(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	spanId := trace.SpanContextFromContext(ctx).SpanID
	if spanId.IsValid() {
		return spanId.String()
	}
	return ""
}

// SetBaggageValue is a convenient function for adding one key-value pair to baggage.
// Note that it uses label.Any to set the key-value pair.
func SetBaggageValue(ctx context.Context, key string, value interface{}) context.Context {
	return NewBaggage(ctx).SetValue(key, value)
}

// SetBaggageMap is a convenient function for adding map key-value pairs to baggage.
// Note that it uses label.Any to set the key-value pair.
func SetBaggageMap(ctx context.Context, data map[string]interface{}) context.Context {
	return NewBaggage(ctx).SetMap(data)
}

// GetBaggageMap retrieves and returns the baggage values as map.
func GetBaggageMap(ctx context.Context) *gmap.StrAnyMap {
	return NewBaggage(ctx).GetMap()
}

// GetBaggageVar retrieves value and returns a *gvar.Var for specified key from baggage.
func GetBaggageVar(ctx context.Context, key string) *gvar.Var {
	return NewBaggage(ctx).GetVar(key)
}
