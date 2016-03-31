package server

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/intervention-engine/fhir/models"
	. "gopkg.in/check.v1"
)

type BindSuite struct {
}

var _ = Suite(&BindSuite{})

func (b *BindSuite) TestJSONBinding(c *C) {
	testBinding(c, "application/json")
}

func (b *BindSuite) TestJSONFHIRBinding(c *C) {
	testBinding(c, "application/json+fhir")
}

func testBinding(c *C, contentType string) {
	data, _ := os.Open("../fixtures/condition.json")

	ctx, _, _ := gin.CreateTestContext()
	ctx.Request, _ = http.NewRequest("POST", "/Condition", data)
	ctx.Request.Header.Add("Content-Type", contentType)

	var condition models.Condition
	FHIRBind(ctx, &condition)

	c.Assert(condition.ResourceType, Equals, "Condition")
	c.Assert(condition.Id, Equals, "8664777288161060797")
	c.Assert(condition.VerificationStatus, Equals, "confirmed")
	c.Assert(condition.Patient, NotNil)
	c.Assert(condition.Patient.Reference, Equals, "https://example.com/base/Patient/4954037118555241963")
	c.Assert(condition.Code, NotNil)
	c.Assert(condition.Code.Text, Equals, "Heart failure")
	c.Assert(condition.Code.Coding, HasLen, 3)
	c.Assert(condition.Code.MatchesCode("http://snomed.info/sct", "10091002"), Equals, true)
	c.Assert(condition.Code.MatchesCode("http://hl7.org/fhir/sid/icd-9", "428.0"), Equals, true)
	c.Assert(condition.Code.MatchesCode("http://hl7.org/fhir/sid/icd-10", "I50.1"), Equals, true)
	c.Assert(condition.OnsetDateTime, NotNil)
	tz, _ := time.LoadLocation("America/New_York")
	c.Assert(condition.OnsetDateTime.Time.Equal(time.Date(2012, time.March, 1, 7, 0, 0, 0, tz)), Equals, true)
	c.Assert(condition.OnsetDateTime.Precision, Equals, models.Precision(models.Timestamp))
}
