package pages

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mbrostami/gcron/internal/db"
	pb "github.com/mbrostami/gcron/internal/grpc"
)

// TaskPage using Page interface
type TaskPage struct {
	db db.DB
}

// NewTaskPage creates new page
func NewTaskPage(db db.DB) *TaskPage {
	return &TaskPage{db: db}
}

// GetRoute url endpoint
func (p *TaskPage) GetRoute() string {
	return "/tasks/:uid"
}

// GetMethods method available for this page
func (p *TaskPage) GetMethods() []string {
	return []string{"GET"}
}

// Handler get page parameters
func (p *TaskPage) Handler(method string, c *gin.Context) Response {
	user, _ := c.Get("user")
	var res Response
	if user == nil {
		res = gin.H{
			"message": "You need to login first!",
		}
		c.Redirect(http.StatusFound, "/login")
		return res
	}
	limit, _ := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 32)
	from, _ := strconv.ParseInt(c.DefaultQuery("from", "0"), 10, 32)

	uid, err := strconv.ParseUint(c.Param("uid"), 10, 32)
	if err != nil {
		c.HTML(500, "error.tmpl", gin.H{
			"message": err.Error(),
		})
	}
	taskCollection := p.db.Get(uint32(uid), int(from), int(limit-1))

	// As all tasks coming by uid has same command
	// Get first one to extract command name
	var command *pb.Task
	for _, task := range taskCollection.Tasks {
		command = task
		break
	}
	res = gin.H{
		"user":    user,
		"command": command,
		"tasks":   taskCollection.Tasks,
	}
	c.HTML(200, "task.tmpl", res)
	return res
}
