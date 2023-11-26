package utils

import (
	// "fmt"
	// "errors"

	uuid "github.com/satori/go.uuid"

	// "net"
	"time"
)

// 获取机器的UID
func GetLocalUUID() string {
	ul := uuid.NewV4()
	return ul.String()
}

func GetNowTimeStamp() int64 {
	currentTime := time.Now().UnixMilli() //获取当前时间戳，单位ms
	return currentTime
}
