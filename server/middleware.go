package server

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func EnableXmlToJsonConversionMiddleware() gin.HandlerFunc {
	converter := NewFhirFormatConverter()
	return func(c *gin.Context) {
		c.Set("FhirFormatConverter", converter)
		c.Next()
	}
}

// AbortNonJSONRequestsMiddleware is middleware that responds to any request that Accepts a Content-Type
// other than JSON (or a JSON flavor) with a 406 Not Acceptable status.
func AbortNonJSONRequestsMiddleware(c *gin.Context) {
	acceptHeader := c.Request.Header.Get("Accept")
	formatOption := c.DefaultQuery("_format", "")
	hasJSON := hasJsonMimeType(acceptHeader, formatOption) > 0 || strings.Contains(acceptHeader, "json") // allowing non-FHIR MIME types as per previous version
	if acceptHeader != "" && !hasJSON && !strings.Contains(acceptHeader, "*/*") {
		c.AbortWithStatus(http.StatusNotAcceptable)
	}
	c.Next()
}

// AbortNonFhirXMLorJSONRequestsMiddleware is middleware that responds to any request that Accepts a Content-Type
// other than FHIR JSON or XML with a 406 Not Acceptable status.
func AbortNonFhirXMLorJSONRequestsMiddleware(c *gin.Context) {
	acceptHeader := c.Request.Header.Get("Accept")
	formatOption := c.DefaultQuery("_format", "")
	hasJSON := hasJsonMimeType(acceptHeader, formatOption)
	hasXML := hasXmlMimeType(acceptHeader, formatOption)
	if acceptHeader != "" && hasXML == 0 && hasJSON == 0 && !strings.Contains(acceptHeader, "*/*") {
		c.AbortWithStatus(http.StatusNotAcceptable)
	}
	if hasXML > hasJSON { // integer comparison so that _format overrides an Accept header
		c.Set("SendXML", true)
	}
	c.Next()
}

func hasJsonMimeType(acceptHeader string, formatOption string) int {
	// _format overrides the Accept header according to the spec
	switch formatOption {
	case "json":
	case "text/json":
	case "application/json":
	case "application/fhir+json":
		return 2
	}
	if strings.Contains(acceptHeader, "application/fhir+json") || strings.Contains(acceptHeader, "application/json+fhir") {
		return 1
	} else {
		return 0
	}
}
func hasXmlMimeType(acceptHeader string, formatOption string) int {
	// _format overrides the Accept header according to the spec
	switch formatOption {
	case "xml":
	case "text/xml":
	case "application/xml":
	case "application/fhir+xml":
		return 2
	}
	if strings.Contains(acceptHeader, "application/fhir+xml") || strings.Contains(acceptHeader, "application/xml+fhir") {
		return 1
	} else {
		return 0
	}
}

// ReadOnlyMiddleware makes the API read-only and responds to any requests that are not
// GET, HEAD, or OPTIONS with a 405 Method Not Allowed error.
func ReadOnlyMiddleware(c *gin.Context) {
	method := c.Request.Method
	switch method {
	// allowed methods:
	case "GET", "HEAD", "OPTIONS":
		c.Next()
	// all other methods:
	default:
		c.AbortWithStatus(http.StatusMethodNotAllowed)
	}
}
