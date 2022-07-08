package model

import "time"

// LogLayout 日志layout
type LogLayout struct {
	Time      time.Time
	UserID    int // 用户
	UserName  string
	Path      string        // 访问路径
	Query     string        // 携带query
	Body      string        // 携带body数据
	IP        string        // ip地址
	UserAgent string        // 代理
	Error     string        // 错误
	Cost      time.Duration // 花费时间
	Source    string        // 来源
}
