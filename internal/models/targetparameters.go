package models

import "net/http"

// TargetParameters contains all the information, later becomes an item of flowStack. 
type TargetParameters struct {
	ApiType string
	Method  string
	URL     string
	Header  http.Header
	Body    []byte
}
