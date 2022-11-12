package driver

import (
	"fmt"
	"sync"
	"task_list/config"
	"task_list/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	gormEngine     *gorm.DB
	gormEngineOnce sync.Once
)

// InitGorm return singleton gorm instance
func InitGorm() *gorm.DB {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci&loc=UTC",
		config.GlobalConfig.MySQLConfig.Username,
		config.GlobalConfig.MySQLConfig.Password,
		config.GlobalConfig.MySQLConfig.Host,
		config.GlobalConfig.MySQLConfig.Port,
		config.GlobalConfig.MySQLConfig.DBName,
	)
	gormEngineOnce.Do(func() {
		gormEngine, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
		if err != nil {
			panic(err)
		}
		err = gormEngine.AutoMigrate(models.Task{})
		if err != nil {
			panic(err)
		}
	})
	return gormEngine

}
