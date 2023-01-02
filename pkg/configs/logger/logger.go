package logger

import (
	"os"
	"strconv"
	"time"

	"github.com/cj2172171/xh-commons/pkg/configs/sysconfigs"
	"github.com/luaxlou/goutils/tools/timeutils"
	"github.com/sirupsen/logrus"
)

func Init() {

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000000",
	})
	logrus.SetOutput(os.Stdout)

	timeutils.ImmediatelyTick(setLevel, 10*time.Second)
}

var cfg map[string]string

func setLevel() {

	cfg = sysconfigs.GetConfig("log")
	if cfg == nil {
		logrus.SetLevel(logrus.InfoLevel)
		return
	}

	level, _ := strconv.Atoi(cfg["level"])
	switch level {
	case 3:
		logrus.SetLevel(logrus.WarnLevel)
	case 4:
		logrus.SetLevel(logrus.InfoLevel)
	case 5:
		logrus.SetLevel(logrus.DebugLevel)
	case 6:
		logrus.SetLevel(logrus.TraceLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

}
