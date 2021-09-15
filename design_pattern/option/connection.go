package option

import "time"

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

const (
	defaultCache   = true
	defaultTimeout = 100
)

type options struct {
	cache   bool
	timeout time.Duration
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithCache(cache bool) Option {
	return optionFunc(func(o *options) {
		o.cache = cache
	})
}

func WithTimeout(timeout time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = timeout
	})
}

func NewConnection(addr string, os ...Option) *Connection {
	ops := &options{
		cache:   defaultCache,
		timeout: defaultTimeout,
	}

	for _, o := range os {
		o.apply(ops)
	}

	return &Connection{
		addr:    addr,
		cache:   ops.cache,
		timeout: ops.timeout,
	}
}
