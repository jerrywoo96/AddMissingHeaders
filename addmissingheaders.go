package AddMissingHeaders

import (
	"context"
	"net/http"
	"strings"
)

// Config holds configuration to be passed to the plugin
type Config struct {
	RequestHeaders  map[string]string `yaml:"requestHeaders,omitempty"`
	ResponseHeaders map[string]string `yaml:"responseHeaders,omitempty"`
}

// CreateConfig populates the Config data object
func CreateConfig() *Config {
	return &Config{}
}

// Plugin holds the necessary components of a Traefik plugin
type Plugin struct {
	requestHeaders  map[string]string
	responseHeaders map[string]string
	name            string
	next            http.Handler
}

// New instantiates and returns the required components used to handle a HTTP request
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Plugin{
		requestHeaders:  config.RequestHeaders,
		responseHeaders: config.ResponseHeaders,
		name:            name,
		next:            next,
	}, nil
}

func (plugin *Plugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	for key, value := range plugin.requestHeaders {
		header := req.Header.Get(strings.TrimSpace(key))
		if len(strings.TrimSpace(header)) == 0 {
			req.Header.Set(strings.TrimSpace(key), strings.TrimSpace(value))
		}
	}

	for key, value := range plugin.responseHeaders {
		header := rw.Header().Get(strings.TrimSpace(key))
		if len(strings.TrimSpace(header)) == 0 {
			rw.Header().Add(strings.TrimSpace(key), strings.TrimSpace(value))
		}
	}

	plugin.next.ServeHTTP(rw, req)

}
