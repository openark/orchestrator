package abstractions

import "strings"

//RequestHeaders represents a collection of request headers
type RequestHeaders struct {
	headers map[string]map[string]struct{}
}

//NewRequestHeaders creates a new RequestHeaders
func NewRequestHeaders() *RequestHeaders {
	return &RequestHeaders{
		headers: make(map[string]map[string]struct{}),
	}
}

func (r *RequestHeaders) normalizeHeaderKey(key string) string {
	return strings.ToLower(strings.Trim(key, " "))
}

type void struct{}

var voidInstance void

//Add adds a new header or append a new value to an existing header
func (r *RequestHeaders) Add(key string, value string, additionalValues ...string) {
	normalizedKey := r.normalizeHeaderKey(key)
	if normalizedKey == "" || value == "" {
		return
	}
	if r.headers == nil {
		r.headers = make(map[string]map[string]struct{})
	}
	if r.headers[normalizedKey] == nil {
		r.headers[normalizedKey] = make(map[string]struct{})
	}
	r.headers[normalizedKey][value] = voidInstance
	for _, v := range additionalValues {
		r.headers[normalizedKey][v] = voidInstance
	}
}

//Get returns the values for the specific header
func (r *RequestHeaders) Get(key string) []string {
	if r.headers == nil {
		return nil
	}
	normalizedKey := r.normalizeHeaderKey(key)
	if r.headers[normalizedKey] == nil {
		return make([]string, 0)
	}
	values := make([]string, 0, len(r.headers[normalizedKey]))
	for k := range r.headers[normalizedKey] {
		values = append(values, k)
	}
	return values
}

//Remove removes the specific header and all its values
func (r *RequestHeaders) Remove(key string) {
	if r.headers == nil {
		return
	}
	normalizedKey := r.normalizeHeaderKey(key)
	delete(r.headers, normalizedKey)
}

//RemoveValue remove the value for the specific header
func (r *RequestHeaders) RemoveValue(key string, value string) {
	if r.headers == nil {
		return
	}
	normalizedKey := r.normalizeHeaderKey(key)
	if r.headers[normalizedKey] == nil {
		return
	}
	delete(r.headers[normalizedKey], value)
	if len(r.headers[normalizedKey]) == 0 {
		delete(r.headers, normalizedKey)
	}
}

//ContainsKey check if the key exists in the headers
func (r *RequestHeaders) ContainsKey(key string) bool {
	if r.headers == nil {
		return false
	}
	normalizedKey := r.normalizeHeaderKey(key)
	return r.headers[normalizedKey] != nil
}

//Clear clear all headers
func (r *RequestHeaders) Clear() {
	r.headers = nil
}

//AddAll adds all headers from the other RequestHeaders
func (r *RequestHeaders) AddAll(other *RequestHeaders) {
	if other == nil || other.headers == nil {
		return
	}
	for k, v := range other.headers {
		for k2 := range v {
			r.Add(k, k2)
		}
	}
}

//ListKeys returns all the keys in the headers
func (r *RequestHeaders) ListKeys() []string {
	if r.headers == nil {
		return make([]string, 0)
	}
	keys := make([]string, 0, len(r.headers))
	for k := range r.headers {
		keys = append(keys, k)
	}
	return keys
}
