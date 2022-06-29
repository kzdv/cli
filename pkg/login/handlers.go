package login

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kzdv/cli/pkg/creds"
)

// This will get set during build
//go:embed client_secret.txt
var CLIENT_SECRET string

const REDIRECT_URI = "http://127.0.0.1:12297/callback"

func HandleAuthorize(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect,
		"https://auth.denartcc.org/oauth/authorize?response_type=code&client_id=kubernetes&scope=openid+full_name+email&redirect_uri="+
			url.QueryEscape(REDIRECT_URI))
}

func HandleCallback(c *gin.Context) {
	data := url.Values{
		"scope":         {"openid", "full_name", "email"},
		"response_type": {"id_token"},
		"client_id":     {"kubernetes"},
		"client_secret": {CLIENT_SECRET},
		"code":          {c.Query("code")},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {REDIRECT_URI},
	}

	tokenReq, err := http.PostForm("https://auth.denartcc.org/oauth/token", data)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "return.tmpl", gin.H{
			"title":   "Error",
			"message": "There was an error while trying to make the token request",
		})
		fmt.Printf("Error making token request: %s", err)
		os.Exit(1)
	}
	defer tokenReq.Body.Close()

	body, err := ioutil.ReadAll(tokenReq.Body)
	if err != nil || tokenReq.StatusCode > 399 {
		c.HTML(http.StatusInternalServerError, "return.tmpl", gin.H{
			"title":   "Error",
			"message": "There was an error while trying to read the token request body",
		})
		fmt.Printf("Error reading token request body: %s %s", err, string(body))
		os.Exit(1)
	}

	tokens := &creds.Creds{}
	err = json.Unmarshal(body, tokens)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "return.tmpl", gin.H{
			"title":   "Error",
			"message": "There was an error while trying to unmarshal the token request body",
		})
		fmt.Printf("Error unmarshaling token request body: %s", err)
		os.Exit(1)
	}

	c.HTML(http.StatusOK, "return.tmpl", gin.H{
		"title":   "Success",
		"message": "You are now logged in. You may close this browser window.",
	})

	server.Chan <- *tokens
}
