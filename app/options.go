package app

import "context"

type Option struct {
	context    context.Context
	configFile string
	cors       bool
}

type CallOption func(option *Option)

func newOption(o ...CallOption) *Option {
	opt := &Option{
		context: context.Background(),
		cors:    true,
	}

	for _, oo := range o {
		oo(opt)
	}
	return opt
}

func WithCors(b bool) CallOption {
	return func(o *Option) {
		o.cors = b
	}
}

func WithContext(ctx context.Context) CallOption {
	return func(o *Option) {
		o.context = ctx
	}
}

func WithConfigFile(configFile string) CallOption {
	return func(o *Option) {
		o.configFile = configFile
	}
}
