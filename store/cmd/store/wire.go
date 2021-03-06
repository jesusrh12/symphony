// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build wireinject

package main

import (
	"context"
	"fmt"

	"github.com/facebookincubator/symphony/pkg/log"
	"github.com/facebookincubator/symphony/pkg/server/metrics"
	"github.com/facebookincubator/symphony/pkg/server/xserver"
	"github.com/facebookincubator/symphony/store/handler"
	"github.com/google/wire"
	"gocloud.dev/blob"
	"gocloud.dev/server/health"
)

func newApplication(ctx context.Context, flags *cliFlags) (*application, func(), error) {
	wire.Build(
		wire.Struct(new(application), "*"),
		wire.FieldsOf(new(*cliFlags),
			"ListenAddress",
			"MetricsAddress",
			"LogConfig",
			"TelemetryConfig",
		),
		log.Provider,
		metrics.Provider,
		xserver.ServiceSet,
		xserver.DefaultViews,
		wire.Value([]health.Checker(nil)),
		provideBucket,
		handler.Provider,
	)
	return nil, nil, nil
}

func provideBucket(ctx context.Context, flags *cliFlags) (*blob.Bucket, func(), error) {
	bucket, err := blob.OpenBucket(ctx, flags.BucketURL.String())
	if err != nil {
		return nil, nil, fmt.Errorf("cannot open blob bucket: %w", err)
	}
	return bucket, func() { _ = bucket.Close() }, nil
}
