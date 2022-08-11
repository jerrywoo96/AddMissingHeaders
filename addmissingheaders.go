package AddMissingHeaders

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"net/http"
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
		if values := req.Header.Values(key); values == nil {
			req.Header.Set(key, value)
		}
	}

	if len(plugin.responseHeaders) == 0 {
		plugin.next.ServeHTTP(rw, req)
		return
	}

	plugin.next.ServeHTTP(newResponseModifier(plugin.responseHeaders, rw), req)
}

type responseModifier struct {
	rw http.ResponseWriter

	responseHeaders map[string]string

	headersSent bool // whether headers have already been sent
	code        int  // status code, must default to 200
}

// modifier can be nil.
func newResponseModifier(responseHeaders map[string]string, w http.ResponseWriter) http.ResponseWriter {
	rm := &responseModifier{
		rw:              w,
		code:            http.StatusOK,
		responseHeaders: responseHeaders,
	}

	return rm
}

func (r *responseModifier) Header() http.Header {
	return r.rw.Header()
}

func (r *responseModifier) WriteHeader(code int) {
	if r.headersSent {
		return
	}
	defer func() {
		r.code = code
		r.headersSent = true
	}()

	for key, value := range r.responseHeaders {
		if values := r.rw.Header().Values(key); values == nil {
			r.rw.Header().Set(key, value)
		}
	}

	r.rw.WriteHeader(code)
}

func (r *responseModifier) Write(b []byte) (int, error) {
	r.WriteHeader(r.code)

	return r.rw.Write(b)
}

// Hijack hijacks the connection.
func (r *responseModifier) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if h, ok := r.rw.(http.Hijacker); ok {
		return h.Hijack()
	}

	return nil, nil, fmt.Errorf("not a hijacker: %T", r.rw)
}

// Flush sends any buffered data to the client.
func (r *responseModifier) Flush() {
	if flusher, ok := r.rw.(http.Flusher); ok {
		flusher.Flush()
	}
}
