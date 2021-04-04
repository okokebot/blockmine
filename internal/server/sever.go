package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/okokebot/blockmine/internal/importid"
)

type Request struct {
	Body string `json:"body"`
}

func Server() {
	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowOrigins: []string{
			"*",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))
	blockissueEngine := engine.Group("/blockissues")
	{
		blockissueEngine.POST("", convertController)
	}
	engine.Run(":8080")
}

func convertController(c *gin.Context) {
	fmt.Println("=========")
	fmt.Println(c)
	fmt.Println("=========")
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request"})
		return
	}

	result := importid.CreateBlockNote(req.Body)
	c.JSON(http.StatusOK, gin.H{
		"body": result,
	})
}
