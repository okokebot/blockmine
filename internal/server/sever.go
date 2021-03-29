package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	// c.ShouldBindJSON(&req)
	// fmt.Println("=========")
	// fmt.Println(req)
	// fmt.Println("=========")

	// c.Bind(&req)
	// fmt.Println("=========")
	// fmt.Println(req)
	// fmt.Println("=========")
	// パラメータ取得
	// title := c.Query("title")
	// category, _ := strconv.Atoi(c.Query("category"))
	// author := c.Query("author")

	// 検索処理
	// var s book.Service
	// p, err := s.Search(title, category, author)≥
	fmt.Println(req.Body)
	result := "req.Bodyでない"
	// // 検索結果を返す
	// if err != nil {
	// 	c.AbortWithStatus(http.StatusNotFound)
	// 	fmt.Println(err)
	// } else {
	// 	c.JSON(http.StatusOK, p)
	// }
	c.JSON(http.StatusOK, gin.H{
		"body": result,
	})
}
