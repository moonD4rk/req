// Package req implements a friendly API over Go's existing net/http library
package req

// Get takes 2 parameters and returns a Response Struct. These two options are:
// 	1. A URL
// 	2. A RqOptions struct
// If you do not intend to use the `RqOptions` you can just pass nil
func Get(url string, ro *RqOptions) (*Response, error) {
	return DoRegularRequest(MethodGet, url, ro)
}

// Put takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A RqOptions struct
// If you do not intend to use the `RqOptions` you can just pass nil
func Put(url string, ro *RqOptions) (*Response, error) {
	return DoRegularRequest(MethodPut, url, ro)
}

// Patch takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A RqOptions struct
// If you do not intend to use the `RqOptions` you can just pass nil
func Patch(url string, ro *RqOptions) (*Response, error) {
	return DoRegularRequest(MethodPatch, url, ro)
}

// Delete takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A RqOptions struct
// If you do not intend to use the `RqOptions` you can just pass nil
func Delete(url string, ro *RqOptions) (*Response, error) {
	return DoRegularRequest(MethodDelete, url, ro)
}

// Post takes 2 parameters and returns a Response channel. These two options are:
// 	1. A URL
// 	2. A RqOptions struct
// If you do not intend to use the `RqOptions` you can just pass nil
func Post(url string, ro *RqOptions) (*Response, error) {
	return DoRegularRequest(MethodPost, url, ro)
}

// Head takes 2 parameters and returns a Response channel. These two options are:
// 	1. A URL
// 	2. A RqOptions struct
// If you do not intend to use the `RqOptions` you can just pass nil
func Head(url string, ro *RqOptions) (*Response, error) {
	return DoRegularRequest(MethodHead, url, ro)
}

// Options takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A RqOptions struct
// If you do not intend to use the `RqOptions` you can just pass nil
func Options(url string, ro *RqOptions) (*Response, error) {
	return DoRegularRequest(MethodOptions, url, ro)
}

// Req takes 3 parameters and returns a Response Struct. These three options are:
//	1. A verb
// 	2. A URL
// 	3. A RqOptions struct
// If you do not intend to use the `RqOptions` you can just pass nil
func Req(verb string, url string, ro *RqOptions) (*Response, error) {
	return DoRegularRequest(verb, url, ro)
}
