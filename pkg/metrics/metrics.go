// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

// Package metrics provides a framework for application instrumentation.
//
// Taken from go-kit metrics package
package metrics

// Counter describes a metric that accumulates values monotonically.
// An example of a counter is the number of received HTTP requests.
type Counter interface {
	With(labelValues ...string) Counter
	Add(delta float64)
}

// Gauge describes a metric that takes specific values over time.
// An example of a gauge is the current depth of a job queue.
type Gauge interface {
	With(labelValues ...string) Gauge
	Set(value float64)
	Add(delta float64)
}

// Histogram describes a metric that takes repeated observations of the same
// kind of thing, and produces a statistical summary of those observations,
// typically expressed as quantiles or buckets. An example of a histogram is
// HTTP request latencies.
type Histogram interface {
	With(labelValues ...string) Histogram
	Observe(value float64)
}
