package middleware

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
)

type InterfaceLog struct {
}

// Format 实现logrus的Format接口，实现自定义日志格式
func (log *InterfaceLog) Format(entry *logrus.Entry) ([]byte, error) {
	var buf *bytes.Buffer
	if entry.Buffer != nil {
		buf = entry.Buffer
	} else {
		buf = &bytes.Buffer{}
	}
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	// 日志记录格式
	formatString := fmt.Sprintf("", timestamp, entry)

	buf.WriteString(formatString)
	return buf.Bytes(), nil
}

//func InterfaceLogger() gin.HandlerFunc {
//	// 自定义Formatter
//	loggerFormatter := func(param gin.LogFormatterParams) string {
//		return FormatLog(param)
//	}
//	// 重定向io.writer
//	writer := LogToFile()
//
//	// 配置LoggerConfig
//	conf := gin.LoggerConfig{
//		Formatter: loggerFormatter,
//		Output:    writer,
//	}
//	// 注册
//	return gin.LoggerWithConfig(conf)
//}
//
//func LogToFile() io.Writer {
//	logfile, err := os.Create("logs/gin_http.log")
//	if err != nil {
//		fmt.Println("Could not create log file")
//	}
//	return io.MultiWriter(logfile)
//}
