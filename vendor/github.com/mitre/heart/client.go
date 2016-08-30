package heart

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/oauth2"

	"github.com/icrowley/fake"
	"github.com/juju/errors"

	"gopkg.in/square/go-jose.v1"
)

// ClientJWT represents the JWT used to authenticate a client to a token
// endpoint as specified in Health Relationship Trust Profile for OAuth 2.0 -
// Section 2.2: http://openid.bitbucket.org/HEART/openid-heart-oauth2.html#rfc.section.2.2
type ClientJWT struct {
	ISS string `json:"iss"`
	SUB string `json:"sub"`
	AUD string `json:"aud"`
	IAT int64  `json:"iat"`
	EXP int64  `json:"exp"`
	JTI string `json:"jti"`
}

// NewClientJWT creates a ClientJWT. ISS and SUB are set to the same thing.
// IAT is set to the current time and EXP is set 60 seconds later.
func NewClientJWT(iss string, aud string) ClientJWT {
	jwt := ClientJWT{}
	jwt.ISS = iss
	jwt.SUB = iss
	jwt.AUD = aud
	now := time.Now()
	jwt.IAT = now.Unix()
	jwt.EXP = jwt.IAT + 60
	jwt.JTI = fake.CharactersN(50)
	return jwt
}

// SignJWT takes a ClientJWT, marshals it into JSON, signs the JSON with the
// JWK provided and then returns the blob as a string.
func SignJWT(jwt ClientJWT, pk jose.JsonWebKey) (string, error) {
	signer, err := jose.NewSigner(jose.RS512, &pk)
	if err != nil {
		return "", errors.Annotate(err, "Couldn't create JWT Signer")
	}
	json, err := json.Marshal(jwt)
	if err != nil {
		return "", errors.Annotate(err, "Couldn't marshal the JWT")
	}
	jws, err := signer.Sign(json)
	if err != nil {
		return "", errors.Annotate(err, "Couldn't sign the JWT")
	}
	return jws.CompactSerialize()
}

// Client represents an OAuth 2.0 client that conforms with the HEART
// profile.
type Client struct {
	ISS         string
	AUD         string
	Endpoint    oauth2.Endpoint
	Scopes      []string
	PrivateKey  jose.JsonWebKey
	RedirectURI string
}

// Exchange swaps an authorization code for a token. This
// exchange is compliant with the HEART profile for requests to the
// token endpoint as defined by: http://openid.bitbucket.org/HEART/openid-heart-oauth2.html#rfc.section.2.2
func (c Client) Exchange(code string) (*oauth2.Token, error) {
	jwt := NewClientJWT(c.ISS, c.AUD)
	clientAssertion, err := SignJWT(jwt, c.PrivateKey)
	if err != nil {
		return nil, err
	}
	values := url.Values{"client_assertion": {clientAssertion},
		"client_assertion_type": {"urn:ietf:params:oauth:client-assertion-type:jwt-bearer"},
		"client_id":             {c.ISS}, "grant_type": {"authorization_code"},
		"code": {code}}
	resp, err := http.PostForm(c.Endpoint.TokenURL, values)
	if err != nil {
		return nil, errors.Annotate(err, "Couldn't connect to the token endpoint")
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	token := &oauth2.Token{}
	err = decoder.Decode(token)
	if err != nil {
		return nil, errors.Annotate(err, "Couldn't decode the token")
	}
	return token, nil
}

// AuthURL provides a URL to redirect the client for authorization
// at the authorization server
func (c Client) AuthURL(state string) string {
	values := url.Values{"client_id": {c.ISS}, "state": {"state"}, "response_type": {"code"}}
	return fmt.Sprintf("%s?%s", c.Endpoint.AuthURL, values.Encode())
}
