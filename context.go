package honey

import (
	"net/http"
	"net/url"
)

// Context includes http request info.
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Multipart      bool
	Form           url.Values
}
