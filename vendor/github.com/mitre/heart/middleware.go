package heart

import (
	"encoding/gob"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/icrowley/fake"
	"github.com/juju/errors"
	"gopkg.in/square/go-jose.v1"
)

// This package initialization function is used to register types with gob
// so that they can be stored in sessions. github.com/gin-gonic/contrib/sessions
// uses gob to write structs out into a cookie.
func init() {
	gob.Register(OpenIDTokenResponse{})
	gob.Register(UserInfo{})
}

// OAuthIntrospectionHandler creates a gin.HandlerFunc that can be used to introspect
// OAuth 2.0 tokens provided in the request. endpoint is the address of the authorization
// server token introspection service. iss is the client id for the introspection client.
// aud is the audience, which should be the identifier for the authorization server. pk
// is the private key for the client, so it can sign a JWT to authenticate to the introspection
// endpoint.
//
// This middleware will abort any requests that do not have an Authorization header. It will
// also halt requests if the provided bearer token is inactive or expired.
//
// If a valid token is provided, the gin.Context is augmented by setting the following variables:
// scopes will be a []string containing all scopes valid for the provided token. subject will be
// an identifier for the user who delegated the authority represented by the token. clientID will
// contain the identifier for the client issuing the request.
func OAuthIntrospectionHandler(endpoint, iss, aud string, pk jose.JsonWebKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.String(http.StatusForbidden, "No Authorization header provided")
			c.Abort()
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		if token == auth {
			c.String(http.StatusForbidden, "Could not find bearer token in Authorization header")
			c.Abort()
			return
		}
		jwt := NewClientJWT(iss, aud)
		clientAssertion, err := SignJWT(jwt, pk)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ir, err := IntrospectToken(endpoint, token, iss, clientAssertion)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if !ir.Active {
			c.String(http.StatusForbidden, "Provided token is no longer active")
			c.Abort()
			return
		}
		c.Set("scopes", ir.SplitScope())
		c.Set("subject", ir.SUB)
		c.Set("clientID", ir.ClientID)
	}
}

// OIDCAuthenticationHandler is a middleware that will check for the presence of a session with
// a UserInfo value set. If it exists, it will assume that the has logged in at some point. It will then
// check the session for a token, which will be an OpenIDTokenResponse. If it has not expired,
// it will set the UserInfo in a UserInfo value on the gin Context.
// If there is no UserInfo value present in the session or if the OpenIDTokenResponse has expired, the
// user will be redirected to the provided redirectURI.
func OIDCAuthenticationHandler(client Client, op *OpenIDProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Let things pass through to the redirect path, as it sets up the session
		if strings.Contains(c.Request.URL.Path, "/redirect") {
			return
		}
		session := sessions.Default(c)
		ui := session.Get("UserInfo")
		if ui != nil {
			tokenInterface := session.Get("OpenIDTokenResponse")
			if tokenInterface != nil {
				token := tokenInterface.(OpenIDTokenResponse)
				if token.Expiration.After(time.Now()) {
					c.Set("UserInfo", session.Get("UserInfo"))
					return
				}
				valid, err := op.Validate(token.IDToken)
				if err != nil || !valid {
					c.Abort()
					c.String(http.StatusForbidden, "Provided IDToken is not valid or could not be validated")
				}
				session.Delete("UserInfo")
				session.Delete("OpenIDTokenResponse")
			}
		}
		state := fake.CharactersN(20)
		session.Set("state", state)
		err := session.Save()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		authURL := op.AuthURL(client, state)
		c.Redirect(http.StatusFound, authURL)
		c.Abort()
	}
}

// RedirectHandler provides a gin.HandlerFunc to process the authentication response from an
// Open ID Provider.
func RedirectHandler(client Client, op *OpenIDProvider, successfulAuthRedirectURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		authError := c.Query("error")
		if authError != "" {
			session.Delete("state")
			err := errors.Unauthorizedf("OP was unable to successfully authenticate your request: %s", authError)
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		serverState := c.Query("state")
		localState := session.Get("state")
		if localState == nil {
			err := errors.NotValidf("Couldn't find the local state or nonce to verify the response from the OP")
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if localState.(string) != serverState {
			c.String(http.StatusForbidden, "Nonce or state did not match")
			c.Abort()
		}

		code := c.Query("code")
		token, err := op.Exchange(code, client)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		session.Set("OpenIDTokenResponse", token)
		ui, err := op.UserInfo(token.AccessToken)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		session.Set("UserInfo", ui)
		session.Delete("nonce")
		err = session.Save()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.Redirect(http.StatusFound, successfulAuthRedirectURL)
	}
}

// LogoutHandler allows users to log out by clearing all
// values from their session
func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.String(http.StatusOK, "You have successfully logged out.")
}
