package req

import "net/http"

// Session allows a user to make use of persistent cookies in between
// HTTP requests
type Session struct {
	// RqOptions is global options
	RqOptions *RqOptions

	// HTTPClient is the client that we will use to request the resources
	HTTPClient *http.Client
}

// NewSession returns a session struct which enables can be used to maintain establish a persistent state with the
// server
// This function will set UseCookieJar to true as that is the purpose of using the session
func NewSession(ro *RqOptions) *Session {
	if ro == nil {
		ro = &RqOptions{}
	}

	ro.UseCookieJar = true

	return &Session{RqOptions: ro, HTTPClient: BuildHTTPClient(*ro)}
}

// Combine session options and request options
// 1. UserAgent
// 2. Host
// 3. Auth
// 4. Headers
func (s *Session) combineRqOptions(ro *RqOptions) *RqOptions {
	if ro == nil {
		ro = &RqOptions{}
	}

	if ro.UserAgent == "" && s.RqOptions.UserAgent != "" {
		ro.UserAgent = s.RqOptions.UserAgent
	}

	if ro.Host == "" && s.RqOptions.Host != "" {
		ro.Host = s.RqOptions.Host
	}

	if ro.Auth == nil && s.RqOptions.Auth != nil {
		ro.Auth = s.RqOptions.Auth
	}

	if len(s.RqOptions.Headers) > 0 || len(ro.Headers) > 0 {
		headers := make(map[string]string)
		for k, v := range s.RqOptions.Headers {
			headers[k] = v
		}
		for k, v := range ro.Headers {
			headers[k] = v
		}
		ro.Headers = headers
	}
	return ro
}

// Get takes 2 parameters and returns a Response Struct. These two options are:
// 	1. A URL
// 	2. A RqOptions struct
// If you do not intend to use the `RqOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Get(url string, ro *RqOptions) (*Response, error) {
	ro = s.combineRqOptions(ro)
	return doSessionRequest("GET", url, ro, s.HTTPClient)
}

// Put takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A RqOptions struct
// If you do not intend to use the `RqOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Put(url string, ro *RqOptions) (*Response, error) {
	ro = s.combineRqOptions(ro)
	return doSessionRequest("PUT", url, ro, s.HTTPClient)
}

// Patch takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A RqOptions struct
// If you do not intend to use the `RqOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Patch(url string, ro *RqOptions) (*Response, error) {
	ro = s.combineRqOptions(ro)
	return doSessionRequest("PATCH", url, ro, s.HTTPClient)
}

// Delete takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A RqOptions struct
// If you do not intend to use the `RqOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Delete(url string, ro *RqOptions) (*Response, error) {
	ro = s.combineRqOptions(ro)
	return doSessionRequest("DELETE", url, ro, s.HTTPClient)
}

// Post takes 2 parameters and returns a Response channel. These two options are:
// 	1. A URL
// 	2. A RqOptions struct
// If you do not intend to use the `RqOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Post(url string, ro *RqOptions) (*Response, error) {
	ro = s.combineRqOptions(ro)
	return doSessionRequest("POST", url, ro, s.HTTPClient)
}

// Head takes 2 parameters and returns a Response channel. These two options are:
// 	1. A URL
// 	2. A RqOptions struct
// If you do not intend to use the `RqOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Head(url string, ro *RqOptions) (*Response, error) {
	ro = s.combineRqOptions(ro)
	return doSessionRequest("HEAD", url, ro, s.HTTPClient)
}

// Options takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A RqOptions struct
// If you do not intend to use the `RqOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Options(url string, ro *RqOptions) (*Response, error) {
	ro = s.combineRqOptions(ro)
	return doSessionRequest("OPTIONS", url, ro, s.HTTPClient)
}

// CloseIdleConnections closes the idle connections that a session client may make use of
func (s *Session) CloseIdleConnections() {
	s.HTTPClient.Transport.(*http.Transport).CloseIdleConnections()
}
