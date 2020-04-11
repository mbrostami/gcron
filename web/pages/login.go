package pages

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginPage using Page interface
type LoginPage struct {
	user string
	pass string
}

// NewLoginPage creates new page
func NewLoginPage(defuser string, defpass string) *LoginPage {
	return &LoginPage{user: defuser, pass: defpass}
}

// GetRoute url endpoint
func (p *LoginPage) GetRoute() string {
	return "/login"
}

// GetMethods method available for this page
func (p *LoginPage) GetMethods() []string {
	return []string{"GET", "POST"}
}

// Handler get page parameters
func (p *LoginPage) Handler(method string, c *gin.Context) Response {
	session := sessions.Default(c)
	user := session.Get("user")
	var res Response
	if user != nil {
		res = gin.H{
			"message": "Already logged in",
		}
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return res
	}

	if method == "POST" {
		username := c.PostForm("username")
		password := c.PostForm("password")
		// Validate form input
		if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
			res = gin.H{
				"error": "Parameters can't be empty",
			}
			c.HTML(http.StatusBadRequest, "login.tmpl", res)
			return res
		}

		// Check for username and password match, usually from a database
		if username != p.user || password != p.pass {
			res = gin.H{
				"error": "Authentication failed",
			}
			c.HTML(http.StatusUnauthorized, "login.tmpl", res)
			return res
		}
		session.Set("user", username) // login
		if err := session.Save(); err != nil {
			res = gin.H{
				"error": "Failed to save session",
			}
			c.HTML(http.StatusInternalServerError, "login.tmpl", res)
			return res
		}
		res = gin.H{
			"message": "Successfully authenticated user",
		}
		c.Redirect(http.StatusFound, "/")
		return res
	}
	res = gin.H{
		"user": user,
	}
	c.HTML(200, "login.tmpl", res)
	return res
}
