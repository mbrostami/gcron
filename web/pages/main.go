package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mbrostami/gcron/internal/db"
)

// MainPage using Page interface
type MainPage struct {
	db db.DB
}

// NewMainPage creates new page
func NewMainPage(db db.DB) *MainPage {
	return &MainPage{db: db}
}

// GetRoute url endpoint
func (p *MainPage) GetRoute() string {
	return "/"
}

// GetMethods method available for this page
func (p *MainPage) GetMethods() []string {
	return []string{"GET"}
}

// Handler get page parameters
func (p *MainPage) Handler(method string, c *gin.Context) Response {
	user, _ := c.Get("user")
	var res Response
	if user == nil {
		res = gin.H{
			"message": "You need to login first!",
		}
		c.Redirect(http.StatusFound, "/login")
		return res
	}
	taskCollection := p.db.GetTasks(0, 100)
	res = gin.H{
		"commands": taskCollection.Tasks,
		"user":     user,
	}
	c.HTML(200, "main.tmpl", res)
	//c.JSON(200, res)
	return res
}
