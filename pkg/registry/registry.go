package registry

import (
	"context"

	"github.com/nobbmaestro/tmux-tether/pkg/config"
)

type ContextKey int

const (
	ConfigKey ContextKey = iota
	ConfigPathKey
)

type Registry struct {
	Context context.Context
}

type Option func(*Registry)

func NewRegistry(opts ...Option) Registry {
	r := Registry{context.Background()}
	for _, opt := range opts {
		opt(&r)
	}
	return r
}

func WithContext(context context.Context) Option {
	return func(r *Registry) {
		r.Context = context
	}
}

func WithConfig(cfg *config.UserConfig) Option {
	return func(r *Registry) {
		r.Context = context.WithValue(r.Context, ConfigKey, cfg)
	}
}

func WithConfigPath(path string) Option {
	return func(r *Registry) {
		r.Context = context.WithValue(r.Context, ConfigPathKey, path)
	}
}

func (r Registry) GetConfig() *config.UserConfig {
	if val, ok := r.Context.Value(ConfigKey).(*config.UserConfig); ok {
		return val
	}
	return nil
}

func (r Registry) GetConfigPath() string {
	if val, ok := r.Context.Value(ConfigPathKey).(string); ok {
		return val
	}
	return ""
}
