package server

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/okokebot/blockmine/internal/importid"
	"github.com/okokebot/blockmine/pkg/redmine"
)

type JsonRequest struct {
	Body string `json:"body"`
}

func Init() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			"POST",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
	}))
	r.POST("/blockissues", postBlockissues)
	r.Run(":8080")
}

func postBlockissues(c *gin.Context) {
	var json JsonRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"body": blockissuesController(json.Body)})
}

func blockissuesController(s string) string {
	c := redmine.NewClient()
	ids := importid.PulloutIssueIdFromReleaseNote(s)
	fmt.Print(ids)
	result := ""
	for _, id := range ids {
		p := c.GetIssue(id)
		c.GetChildrenInfo(p)
		result += p.CreateReleaseBlock(*c)
	}
	fmt.Println(result)
	return result
}
