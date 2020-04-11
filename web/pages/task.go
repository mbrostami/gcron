package pages

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mbrostami/gcron/internal/db"
	pb "github.com/mbrostami/gcron/internal/grpc"
	"github.com/mbrostami/gcron/pkg/formatters"
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
	command := taskCollection.Tasks[0]
	var xAxis []string
	var yAxis []float64
	var task *pb.Task
	// Make sorted list
	for i := len(taskCollection.Tasks) - 1; i >= 0; i-- {
		task = taskCollection.Tasks[i]
		startTimeUTC := time.Unix(task.StartTime.Seconds, 0)
		xAxis = append(xAxis, startTimeUTC.Format("15:04:05"))
		floatVal, _ := formatters.GetTimestampDiffMS(task.StartTime, task.EndTime)
		yAxis = append(yAxis, floatVal)
	}
	res = gin.H{
		"user":    user,
		"command": command,
		"xAxis":   xAxis,
		"yAxis":   yAxis,
		"tasks":   taskCollection.Tasks,
	}
	c.HTML(200, "task.tmpl", res)
	return res
}
