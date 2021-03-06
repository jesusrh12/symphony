// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import (
	"encoding/base64"
	"fmt"
	"math"
	"time"

	"go.opencensus.io/trace"
	"go.uber.org/zap/zapcore"
)

type (
	spanCore struct {
		zapcore.LevelEnabler
		span  *trace.Span
		attrs []trace.Attribute
	}

	attributes []trace.Attribute
)

// NewSpanCore writes log entries to trace span as attributes.
func NewSpanCore(span *trace.Span, enabler zapcore.LevelEnabler) zapcore.Core {
	return spanCore{
		span:         span,
		LevelEnabler: enabler,
	}
}

func (sc spanCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if sc.Enabled(ent.Level) && sc.span.IsRecordingEvents() {
		return ce.AddCore(ent, sc)
	}
	return ce
}

func (sc spanCore) With(fields []zapcore.Field) zapcore.Core {
	attrs := make(attributes, len(sc.attrs), len(sc.attrs)+len(fields))
	copy(attrs, sc.attrs)
	for _, field := range fields {
		field.AddTo(&attrs)
	}
	sc.attrs = attrs
	return sc
}

func (sc spanCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	attrs := make(attributes, len(sc.attrs), len(sc.attrs)+len(fields)+2)
	copy(attrs, sc.attrs)
	for _, field := range fields {
		field.AddTo(&attrs)
	}
	attrs.AddString("level", ent.Level.String())
	if ent.Stack != "" {
		attrs.AddString("stack", ent.Stack)
	}
	if ent.Level >= zapcore.DPanicLevel {
		sc.span.AddAttributes(trace.BoolAttribute("error", true))
	}
	sc.span.Annotate(attrs, ent.Message)
	return nil
}

func (spanCore) Sync() error {
	return nil
}

func (attrs *attributes) AddBool(key string, value bool) {
	*attrs = append(*attrs, trace.BoolAttribute(key, value))
}

func (attrs *attributes) AddFloat32(key string, value float32) {
	attrs.AddUint32(key, math.Float32bits(value))
}

func (attrs *attributes) AddFloat64(key string, value float64) {
	attrs.AddUint64(key, math.Float64bits(value))
}

func (attrs *attributes) AddInt(key string, value int) {
	attrs.AddInt64(key, int64(value))
}

func (attrs *attributes) AddInt8(key string, value int8) {
	attrs.AddInt64(key, int64(value))
}

func (attrs *attributes) AddInt16(key string, value int16) {
	attrs.AddInt64(key, int64(value))
}

func (attrs *attributes) AddInt32(key string, value int32) {
	attrs.AddInt64(key, int64(value))
}

func (attrs *attributes) AddInt64(key string, value int64) {
	*attrs = append(*attrs, trace.Int64Attribute(key, value))
}

func (attrs *attributes) AddUintptr(key string, value uintptr) {
	attrs.AddInt64(key, int64(value))
}

func (attrs *attributes) AddUint(key string, value uint) {
	attrs.AddInt64(key, int64(value))
}

func (attrs *attributes) AddUint8(key string, value uint8) {
	attrs.AddInt64(key, int64(value))
}

func (attrs *attributes) AddUint16(key string, value uint16) {
	attrs.AddInt64(key, int64(value))
}

func (attrs *attributes) AddUint32(key string, value uint32) {
	attrs.AddInt64(key, int64(value))
}

func (attrs *attributes) AddUint64(key string, value uint64) {
	attrs.AddInt64(key, int64(value))
}

func (attrs *attributes) AddComplex64(key string, value complex64) {
	attrs.AddString(key, fmt.Sprint(value))
}

func (attrs *attributes) AddComplex128(key string, value complex128) {
	attrs.AddString(key, fmt.Sprint(value))
}

func (attrs *attributes) AddDuration(key string, value time.Duration) {
	attrs.AddString(key, value.String())
}

func (attrs *attributes) AddTime(key string, value time.Time) {
	attrs.AddString(key, value.Format(time.RFC3339))
}

func (attrs *attributes) AddString(key, value string) {
	*attrs = append(*attrs, trace.StringAttribute(key, value))
}

func (attrs *attributes) AddBinary(key string, value []byte) {
	attrs.AddString(key, base64.StdEncoding.EncodeToString(value))
}

func (attrs *attributes) AddByteString(key string, value []byte) {
	attrs.AddString(key, string(value))
}

func (attrs *attributes) AddReflected(key string, value interface{}) error {
	attrs.AddString(key, fmt.Sprintf("%v", value))
	return nil
}

func (attrs *attributes) AddArray(key string, marshaler zapcore.ArrayMarshaler) error {
	encoder := zapcore.NewMapObjectEncoder()
	if err := encoder.AddArray(key, marshaler); err != nil {
		return err
	}
	attrs.AddString(key, fmt.Sprint(encoder.Fields[key]))
	return nil
}

func (attrs *attributes) AddObject(key string, marshaler zapcore.ObjectMarshaler) error {
	encoder := zapcore.NewMapObjectEncoder()
	if err := encoder.AddObject(key, marshaler); err != nil {
		return err
	}
	return attrs.AddReflected(key, encoder.Fields[key])
}

func (attributes) OpenNamespace(string) {}
