package sysconfigs

import (
	"testing"

	"github.com/cj2172171/xh-commons/pkg/dbs"
	"github.com/luaxlou/goutils/tools/logutils"
)

func TestGetConfig(t *testing.T) {
	dbs.Init()
	Init()

	c :=GetConfig("message")

	logutils.PrintObj(c)
	c1 :=GetConfig("not exist key")

	logutils.PrintObj(c1)
}
