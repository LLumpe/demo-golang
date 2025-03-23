package mysql

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbName"`
}

// Init 初始化数据库连接.
func Init() {

	dbConfig, err := getConf()
	if err != nil {
		logrus.Errorf("get DB conf failed, err: %v", err)
		panic(err)
	}

	dsn := getDSN(dbConfig)
	logrus.Info(dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Errorf("connect mysql failed, err: %v", err)
		panic(err)
	}
}

// GetDB 获取DB连接。
func GetDB() *gorm.DB {
	return db
}

// getDSN 获取db dsn.
func getDSN(config *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.DBName)
}

// getConf 读取数据库配置.
func getConf() (*DBConfig, error) {
	dataBytes, err := os.ReadFile("./conf/db.yml")
	if err != nil {
		logrus.Errorf("read file failed, err：%v", err)
		return nil, err
	}
	config := &DBConfig{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		logrus.Errorf("parse yml file failed, err：%v", err)
		return nil, err
	}
	return config, nil
}
