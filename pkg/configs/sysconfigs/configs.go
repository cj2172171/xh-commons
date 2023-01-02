package sysconfigs

import (
	"sync"
	"time"

	"github.com/cj2172171/xh-commons/pkg/dbs"
	"github.com/luaxlou/goutils/tools/timeutils"
)

type SysConfig struct {
	Id int64 `gorm:"primary_key"`

	GroupName string `gorm:"unique_index:uk_key"`
	KeyName   string `gorm:"unique_index:uk_key"`
	Value     string
	UpdatedAt time.Time
}

type ConfigMap struct {
	UpdatedAt time.Time

	Data map[string]string
}

var storeMap sync.Map

func getAllConfigs() (cms map[string]ConfigMap) {

	cms = make(map[string]ConfigMap, 0)

	var cs []SysConfig
	dbs.DBRead().Find(&cs)

	for _, c := range cs {

		if _, ok := cms[c.GroupName]; !ok {
			cms[c.GroupName] = ConfigMap{
				UpdatedAt: c.UpdatedAt,
				Data:      make(map[string]string, 0),
			}
		}

		if c.UpdatedAt.After(cms[c.GroupName].UpdatedAt) {

			g := cms[c.GroupName]
			g.UpdatedAt = c.UpdatedAt
			cms[c.GroupName] = g
		}

		cms[c.GroupName].Data[c.KeyName] = c.Value

	}

	return

}

func syncFromDB() {

	cs := getAllConfigs()

	if len(cs) == 0 {
		return
	}

	for groupName, group := range cs {
		if actual, ok := storeMap.LoadOrStore(groupName, group); ok {
			if !actual.(ConfigMap).UpdatedAt.Equal(group.UpdatedAt) {
				storeMap.Store(groupName, group)
			}
		}
	}

}

func Init() {

	dbs.DB().AutoMigrate(&SysConfig{})

	timeutils.ImmediatelyTick(syncFromDB, time.Second*10)

}

func GetConfig(keyName string) map[string]string {

	data, ok := storeMap.Load(keyName)

	if !ok {
		return nil
	}

	return data.(ConfigMap).Data

}
