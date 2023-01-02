package dbs

import (
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/luaxlou/goutils/mysqldb"
	"github.com/luaxlou/goutils/redisdb"
	"gorm.io/gorm"
)

func Init() {

	dbDSN := os.Getenv("DB_DSN")
	dbDSNRead := os.Getenv("DB_DSN_READ")

	redisHost := os.Getenv("REDIS_HOST")
	redisPass := os.Getenv("REDIS_PASS")

	if redisHost != "" {
		redisdb.Init(redisHost, redisPass, 0)

	}

	if dbDSNRead == "" {

		dbDSNRead = dbDSN

	}
	mysqldb.Add("vsiper", dbDSN)
	mysqldb.Add("vsiper-read", dbDSNRead)

}

func DB() *gorm.DB {
	return mysqldb.GetInstance("vsiper")
}

//##读库实例
func DBRead() *gorm.DB {
	return mysqldb.GetInstance("vsiper-read")
}

func RedisDB() *redis.Client {
	return redisdb.DB()
}

func Close() {

	mysqldb.Close()
}
