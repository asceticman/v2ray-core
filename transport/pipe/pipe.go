package pipe

import (
	"v2ray.com/core/common/platform"
	"v2ray.com/core/common/signal"
)

type Option func(*pipe)

func WithoutSizeLimit() Option {
	return func(p *pipe) {
		p.limit = -1
	}
}

func WithSizeLimit(limit int32) Option {
	return func(p *pipe) {
		p.limit = limit
	}
}

func New(opts ...Option) (*Reader, *Writer) {
	p := &pipe{
		limit:       defaultLimit,
		readSignal:  signal.NewNotifier(),
		writeSignal: signal.NewNotifier(),
	}

	for _, opt := range opts {
		opt(p)
	}

	return &Reader{
			pipe: p,
		}, &Writer{
			pipe: p,
		}
}

var defaultLimit int32 = 10 * 1024 * 1024

func init() {
	const raySizeEnvKey = "v2ray.ray.buffer.size"
	size := platform.EnvFlag{
		Name:    raySizeEnvKey,
		AltName: platform.NormalizeEnvName(raySizeEnvKey),
	}.GetValueAsInt(10)
	defaultLimit = int32(size) * 1024 * 1024
}
