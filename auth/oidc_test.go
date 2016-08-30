package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"golang.org/x/oauth2"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mitre/heart"
	"github.com/pebbe/util"
	. "gopkg.in/check.v1"
)

type OIDCSuite struct {
	OIDCServer *httptest.Server
	RP         *gin.Engine
}

var _ = Suite(&OIDCSuite{})

func (o *OIDCSuite) TestUnathorized(c *C) {
	r, err := http.NewRequest("GET", "/", nil)
	c.Assert(err, IsNil)
	rw := httptest.NewRecorder()
	o.RP.ServeHTTP(rw, r)
	c.Assert(http.StatusFound, Equals, rw.Code)
}

func (o *OIDCSuite) TestExchange(c *C) {
	r, err := http.NewRequest("GET", "/", nil)
	util.CheckErr(err)
	rw := httptest.NewRecorder()
	o.RP.ServeHTTP(rw, r)
	c.Assert(http.StatusFound, Equals, rw.Code)
	result := rw.Result()
	cookies := result.Cookies()
	redirectURL := result.Header.Get("Location")
	c.Assert(redirectURL, Not(Equals), "")
	parsedURL, err := url.Parse(redirectURL)
	util.CheckErr(err)
	state := parsedURL.Query().Get("state")
	code := "1234"
	exchangeURL := fmt.Sprintf("/redirect?state=%s&code=%s", state, code)
	exchangeRequest, err := http.NewRequest("GET", exchangeURL, nil)
	util.CheckErr(err)
	exchangeRequest.AddCookie(cookies[0])
	exchangeRecorder := httptest.NewRecorder()
	o.RP.ServeHTTP(exchangeRecorder, exchangeRequest)
	c.Assert(exchangeRecorder.Code, Equals, http.StatusFound)
	c.Assert("/", Equals, exchangeRecorder.Result().Header.Get("Location"))
	authorizedRequest, err := http.NewRequest("GET", "/", nil)
	util.CheckErr(err)
	authorizedRequest.AddCookie(exchangeRecorder.Result().Cookies()[0])
	authorizedRecorder := httptest.NewRecorder()
	o.RP.ServeHTTP(authorizedRecorder, authorizedRequest)
	c.Assert(authorizedRecorder.Code, Equals, http.StatusOK)

}

type returnToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (o *OIDCSuite) SetUpSuite(c *C) {
	accessToken := "access"
	code := "1234"
	e := gin.New()
	e.GET("userInfo", func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		expectedAuthHeader := fmt.Sprintf("Bearer %s", accessToken)
		c.Assert(expectedAuthHeader, Equals, authHeader)
		ctx.JSON(http.StatusOK, &heart.UserInfo{SUB: "foo@bar", Name: "Test User"})
	})
	e.POST("token", func(ctx *gin.Context) {
		providedCode := ctx.PostForm("code")
		c.Assert(code, Equals, providedCode)
		ctx.JSON(http.StatusOK, &returnToken{AccessToken: accessToken, ExpiresIn: 3600})
	})
	server := httptest.NewServer(e)
	o.OIDCServer = server

	e = gin.New()
	store := sessions.NewCookieStore([]byte("cookie"))
	e.Use(sessions.Sessions("mysession", store))

	noop := func(ctx *gin.Context) {
		ui, _ := ctx.Get("UserInfo")
		c.Assert("Test User", Equals, ui.(heart.UserInfo).Name)
		ctx.String(http.StatusOK, "Hello")
	}
	config := oauth2.Config{ClientID: "1234", ClientSecret: "secret", Endpoint: oauth2.Endpoint{
		AuthURL:  "http://doesntmatter.com",
		TokenURL: server.URL + "/token",
	}}

	e.GET("/", OIDCAuthenticationHandler(config), noop)
	e.GET("/redirect", RedirectHandler(config, "/", server.URL+"/userInfo"))

	o.RP = e
}

func (o *OIDCSuite) TearDownSuite(c *C) {
	o.OIDCServer.Close()
}
