// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xserver_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/facebookincubator/symphony/pkg/server/xserver"
	"github.com/stretchr/testify/require"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"gocloud.dev/server/requestlog"
)

func TestLoggerMiddleware(t *testing.T) {
	core, o := observer.New(zap.InfoLevel)
	logger := xserver.NewZapLogger(zap.New(core))
	ctx, span := trace.StartSpan(context.Background(), "test")
	defer span.End()

	req := httptest.NewRequest(http.MethodPost, "/foo/bar", nil).WithContext(ctx)
	rec := httptest.NewRecorder()
	requestlog.NewHandler(logger, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "request error", http.StatusInsufficientStorage)
	})).ServeHTTP(rec, req)

	fields := o.TakeAll()
	require.Len(t, fields, 1)
	require.Equal(t, "http request", fields[0].Message)
	require.Equal(t, zap.InfoLevel, fields[0].Level)
	m := fields[0].ContextMap()
	require.Equal(t, http.MethodPost, m["method"])
	require.EqualValues(t, http.StatusInsufficientStorage, m["status"])
	require.Equal(t, "/foo/bar", m["url"])
	require.Equal(t, span.SpanContext().TraceID.String(), m["trace_id"])
	require.Equal(t, span.SpanContext().SpanID.String(), m["span_id"])
}
