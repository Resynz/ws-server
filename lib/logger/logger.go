/**
 * @Author: Resynz
 * @Date: 2021/5/26 18:11
 */
package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

var (
	Logger = logrus.New()
)

func InitLogger(logPath, logName, level string) error {
	name := path.Join(logPath, logName)
	src, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	Logger.SetOutput(src)

	l, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	Logger.SetLevel(l)
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       false,
	})
	return nil
}
