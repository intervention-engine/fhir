package heart

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"gopkg.in/square/go-jose.v1"
)

func SetUpRoutes(jwkPath, clientID, opURL, serverURL, cookieSecret string, engine *gin.Engine) error {
	jwkJSON, err := os.Open(jwkPath)
	if err != nil {
		return errors.NewNotFound(err, "Couldn't open the RP JWK file")
	}

	jwkBytes, err := ioutil.ReadAll(jwkJSON)
	if err != nil {
		return errors.Annotate(err, "Couldn't read the RP JWK file")
	}

	jwk := jose.JsonWebKey{}
	json.Unmarshal(jwkBytes, &jwk)

	// Set up the HEART Compliant OAuth 2.0 client that will be used by the OIDC RP
	client := Client{
		ISS:         clientID,
		AUD:         opURL + "/",
		RedirectURI: serverURL + "/redirect",
		PrivateKey:  jwk,
	}

	// Set up the new OpenID Provider
	// This will use discovery to find out the authentication, token and user info endpoints
	provider, err := NewOpenIDProvider(opURL)
	if err != nil {
		return errors.Annotate(err, "Couldn't connect to the OIDC server")
	}

	// Pull down the public key for the provider
	err = provider.FetchKey()
	if err != nil {
		return errors.Annotate(err, "Couldn't fetch the OIDC server's public key")
	}

	// Set up sessions so we can keep track of the logged in user
	store := sessions.NewCookieStore([]byte(cookieSecret))
	engine.Use(sessions.Sessions("mysession", store))

	// The OIDCAuthenticationHandler is set up before the IndexHandler in the handler function
	// chain. It will check to see if the user is logged in based on their session. If they are not
	// the user will be redirected to the authentication endpoint at the OP.
	oidcHandler := OIDCAuthenticationHandler(client, provider)
	oauthHandler := OAuthIntrospectionHandler(provider.Config.IntrospectionEndpoint, client.ISS, client.AUD, client.PrivateKey)
	engine.Use(func(c *gin.Context) {
		if c.Request.Header.Get("Authorization") != "" {
			oauthHandler(c)
		} else {
			oidcHandler(c)
		}
	})

	// This handler is to take the redirect from the OP when the user logs in. It will
	// then fetch information about the user by hitting the user info endpoint and put
	// that in the session. Lastly, this handler is set up to redirect the user back
	// to the root.
	engine.GET("/redirect", RedirectHandler(client, provider, serverURL))
	engine.GET("/logout", LogoutHandler)
	return nil
}
