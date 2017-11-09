// Package log provides a structured logger.
//
// Taken from go-kit log package
package log

import (
	"bytes"
	"io"
	"sync"

	"github.com/go-logfmt/logfmt"
)

// Logger is the fundamental interface for all log operations. Log creates a
// log event from keyvals, a variadic sequence of alternating keys and values.
// Implementations must be safe for concurrent use by multiple goroutines. In
// particular, any implementation of Logger that appends to keyvals or
// modifies or retains any of its elements must make a copy first.
type Logger interface {
	Log(keyvals ...interface{}) error
}

func NewSyncWriter(w io.Writer) io.Writer {
	return &syncWriter{Writer: w}
}

// syncWriter synchronizes concurrent writes to an io.Writer.
type syncWriter struct {
	sync.Mutex
	io.Writer
}

// Write writes p to the underlying io.Writer. If another write is already in
// progress, the calling goroutine blocks until the syncWriter is available.
func (w *syncWriter) Write(p []byte) (n int, err error) {
	w.Lock()
	n, err = w.Writer.Write(p)
	w.Unlock()
	return n, err
}

type logfmtEncoder struct {
	*logfmt.Encoder
	buf bytes.Buffer
}

func (l *logfmtEncoder) Reset() {
	l.Encoder.Reset()
	l.buf.Reset()
}

var logfmtEncoderPool = sync.Pool{
	New: func() interface{} {
		var enc logfmtEncoder
		enc.Encoder = logfmt.NewEncoder(&enc.buf)
		return &enc
	},
}

type logfmtLogger struct {
	w io.Writer
}

// NewLogfmtLogger returns a logger that encodes keyvals to the Writer in
// logfmt format. Each log event produces no more than one call to w.Write.
// The passed Writer must be safe for concurrent use by multiple goroutines if
// the returned Logger will be used concurrently.
func NewLogfmtLogger(w io.Writer) Logger {
	return &logfmtLogger{w}
}

func (l logfmtLogger) Log(keyvals ...interface{}) error {
	enc := logfmtEncoderPool.Get().(*logfmtEncoder)
	enc.Reset()
	defer logfmtEncoderPool.Put(enc)

	if err := enc.EncodeKeyvals(keyvals...); err != nil {
		return err
	}

	// Add newline to the end of the buffer
	if err := enc.EndRecord(); err != nil {
		return err
	}

	// The Logger interface requires implementations to be safe for concurrent
	// use by multiple goroutines. For this implementation that means making
	// only one call to l.w.Write() for each call to Log.
	if _, err := l.w.Write(enc.buf.Bytes()); err != nil {
		return err
	}
	return nil
}
