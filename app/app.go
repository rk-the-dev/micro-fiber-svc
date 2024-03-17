package app

import (
	"github.com/rk-the-dev/micro-fiber-svc/helpers/cache"
	"github.com/rk-the-dev/micro-fiber-svc/helpers/config"
	"github.com/rk-the-dev/micro-fiber-svc/helpers/db/ormclient"
	"github.com/rk-the-dev/micro-fiber-svc/helpers/loghelper"
	"github.com/rk-the-dev/micro-fiber-svc/helpers/securityhelper"
	"github.com/sirupsen/logrus"
)

const (
	DB_MY_SQL = "mysql"
	DB_MONGO  = "mongo"
	DB_PostG  = "postgres"
	DB_SQLLIT = "sqllit"
)

var (
	Logger         *logrus.Logger
	Securityhelper *securityhelper.Securityhelper
	MemCache       *cache.CacheHelper
	ConfigHelper   *config.ConfigHelper
	ORMHelper      *ormclient.ORMHelper
)

func InitHelpers() {
	ConfigHelper = config.New()
	ConfigHelper.Load()
	Logger = loghelper.InitLogger()
	Securityhelper = securityhelper.Getsecurityhelper()
	MemCache = cache.NewCacheHelper(-1, -1)
	ORMHelper = ormclient.NewORMHelper()
}
