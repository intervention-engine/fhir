package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mitre/heart"
	"github.com/pebbe/util"

	. "gopkg.in/check.v1"
)

type OAuthSuite struct {
}

var _ = Suite(&OAuthSuite{})

func (o *OAuthSuite) TestIntrospection(c *C) {
	server := httptest.NewServer(mockIntrospectionEndpoint(c))
	defer server.Close()
	rr := oauthRequest("my_client_id", "sekret", server.URL, c)
	c.Assert(rr.Code, Equals, http.StatusOK)
}

func mockIntrospectionEndpoint(c *C) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Assert(r.Method, Equals, "POST")
		clientID := r.FormValue("client_id")
		c.Assert(clientID, Equals, "my_client_id")
		token := r.FormValue("token")
		c.Assert(token, Equals, "foo")
		ir := &heart.IntrospectionResponse{Active: true, Scope: "foo bar", EXP: time.Now().Unix(), SUB: "steve", ClientID: "heart-watch"}
		encoder := json.NewEncoder(w)
		encoder.Encode(ir)
	})
}

func oauthRequest(clientID, clientSecret, endpoint string, c *C) *httptest.ResponseRecorder {
	r, err := http.NewRequest("GET", "/", nil)
	util.CheckErr(err)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer foo")

	e := gin.New()
	rw := httptest.NewRecorder()
	noop := func(ctx *gin.Context) {
		scopes, _ := ctx.Get("scopes")
		c.Assert("bar", Equals, scopes.([]string)[1])
		ctx.String(http.StatusOK, "Hello")
	}

	e.GET("/", OAuthIntrospectionHandler(clientID, clientSecret, endpoint), noop)

	e.ServeHTTP(rw, r)
	return rw
}
