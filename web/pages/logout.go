package pages

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LogoutPage using Page interface
type LogoutPage struct {
}

// NewLogoutPage creates new page
func NewLogoutPage() *LogoutPage {
	return &LogoutPage{}
}

// GetRoute url endpoint
func (p *LogoutPage) GetRoute() string {
	return "/logout"
}

// GetMethods method available for this page
func (p *LogoutPage) GetMethods() []string {
	return []string{"GET"}
}

// Handler get page parameters
func (p *LogoutPage) Handler(method string, c *gin.Context) Response {
	session := sessions.Default(c)
	user := session.Get("user")
	var res Response
	if user != nil {
		session.Delete("user")
		if err := session.Save(); err != nil {
			res = gin.H{
				"message": "Failed to save session",
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
			return res
		}
	}
	res = gin.H{
		"message": "Successfully logged out",
	}
	c.Redirect(http.StatusFound, "/login")
	return res
}
