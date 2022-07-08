package router

import (
	"battery_monitoring/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
	"time"
)

func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		var body []byte
		body, _ = c.GetRawData()
		// 将原body塞回去
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		c.Next()

		cost := time.Since(start)

		// log layout
		layout := model.LogLayout{
			Time:      time.Now(),
			Path:      path,
			Query:     query,
			Body:      string(body),
			IP:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Error:     strings.TrimRight(c.Errors.ByType(gin.ErrorTypePrivate).String(), "\n"),
			Cost:      cost,
			Source:    "asr",
		}
		v, _ := json.Marshal(layout)
		fmt.Println(string(v))
	}
}
