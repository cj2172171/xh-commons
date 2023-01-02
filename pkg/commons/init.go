package commons

import (
	"github.com/cj2172171/xh-commons/pkg/configs/envtools"
	"github.com/cj2172171/xh-commons/pkg/configs/logger"
	"github.com/cj2172171/xh-commons/pkg/configs/sysconfigs"
	"github.com/cj2172171/xh-commons/pkg/dbs"
)

// 最小化系统需求，一些基础组件
func Init() {

	envtools.Init()

	dbs.Init()

	sysconfigs.Init()

	logger.Init()
}

func Freeze() {

	dbs.Close()

}
