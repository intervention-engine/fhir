package server

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func FHIRBind(c *gin.Context, obj interface{}) error {
	if c.Request.Method == "GET" {
		return c.BindWith(obj, binding.Form)
	}
	contentType := c.ContentType()

	if strings.Contains(contentType, "json") {
		return c.BindJSON(obj)
	}

	if strings.Contains(contentType, "application/fhir+xml") || strings.Contains(contentType, "application/xml+fhir") {
		converterInt, enabled := c.Get("FhirFormatConverter")
		if enabled {
			converter := converterInt.(*FhirFormatConverter)
			return converter.ReadXML(c.Request.Body, obj)
		}
	}

	return c.Bind(obj)
}
