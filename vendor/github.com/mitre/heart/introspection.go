package heart

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/juju/errors"
)

// IntrospectionResponse is the response provided by a HEART compliant OAuth 2.0 token introspection
// endpoint as defined at: http://openid.bitbucket.org/HEART/openid-heart-oauth2.html#rfc.section.4.4
type IntrospectionResponse struct {
	Active   bool   `json:"active"`
	Scope    string `json:"scope"`
	EXP      int64  `json:"exp"`
	SUB      string `json:"sub"`
	ClientID string `json:"client_id"`
}

// ExpirationTime converts the EXP value from seconds since the epoch into
// a Time struct
func (ir IntrospectionResponse) ExpirationTime() time.Time {
	return time.Unix(ir.EXP, 0)
}

// SplitScope is a convenience function to split the scope value on spaces
// to get a list of scopes for the provided token
func (ir IntrospectionResponse) SplitScope() []string {
	return strings.Split(ir.Scope, " ")
}

// IntrospectToken provides a way to contact the introspection endpoint specified in the
// endpoint parameter. Token is the opaque token provided by a client attempting to access
// a resource. clientID is the identifier for the introspection client, not the if for the client
// attempting to access the resource. clientAssertion is the signed JWT that will be used to assert
// the identity of this introspection client.
func IntrospectToken(endpoint, token, clientID, clientAssertion string) (IntrospectionResponse, error) {
	values := url.Values{"client_assertion": {clientAssertion},
		"client_assertion_type": {"urn:ietf:params:oauth:client-assertion-type:jwt-bearer"},
		"client_id":             {clientID}, "token": {token}}
	resp, err := http.PostForm(endpoint, values)
	if err != nil {
		return IntrospectionResponse{}, errors.Annotate(err, "Couldn't connect to the introspection endpoint")
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	ir := IntrospectionResponse{}
	err = decoder.Decode(&ir)
	if err != nil {
		return IntrospectionResponse{}, errors.Annotate(err, "Couldn't decode the introspection response")
	}
	return ir, nil
}
