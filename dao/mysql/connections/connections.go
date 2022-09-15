package connections

import (
	"database/sql"
	"fmt"
	"ginDemo/config"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

const (
	Location = "Asia/Shanghai"
	HOST = "rm-2zezcvai82ylssok4.mysql.rds.aliyuncs.com"
	USER = "al_joke_ro"
	PWD = "io1LNDficKPrWw4X"
	DB = "joke"
	MaxOpen = 40
	MaxIdle = 40
)

func MysqlConnectionTest() *sql.DB {
	mysqlConfig := mysql.NewConfig()
	location, err := time.LoadLocation(Location)
	if err != nil {
		log.Panic("load location error", "error:", err.Error())
		return nil
	}
	{
		mysqlConfig.ParseTime = true
		mysqlConfig.ClientFoundRows = true
		mysqlConfig.InterpolateParams = true
		mysqlConfig.Net = "tcp"
		mysqlConfig.Loc = location
		mysqlConfig.Timeout = time.Second * 5
		mysqlConfig.Collation = "utf8mb4_bin"
	}
	mysqlConfig.User = USER
	mysqlConfig.Passwd = PWD
	mysqlConfig.DBName = DB
	mysqlConfig.Addr = HOST
	dsn := mysqlConfig.FormatDSN()
	fmt.Println("mysqlDSN", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic("open mysql db error", "error:", err.Error())
		return nil
	}
	db.SetMaxOpenConns(MaxOpen)
	db.SetMaxIdleConns(MaxIdle)
	db.SetConnMaxLifetime(time.Minute)
	if err := db.Ping(); err != nil {
		log.Panic("mysql ping error", "error:", err.Error())
		return nil
	}
	return db
}


func Mc() *sql.DB {
	mdbConfig := mysql.NewConfig()
	location, err := time.LoadLocation(Location)
	if err != nil {
		log.Panic("load location error", "info: ", err)
		return nil
	}
	{
		mdbConfig.Loc = location
		mdbConfig.Net = "tcp"
	}
	mdbConfig.User = USER
	mdbConfig.Addr = HOST
	mdbConfig.Passwd = PWD
	mdbConfig.DBName = DB
	dsn := mdbConfig.FormatDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic("open mysql error", "info: ", err)
		return nil
	}
	if err := db.Ping(); err != nil {
		log.Panic("ping mysql db error", "info: ", err)
		return nil
	}
	return db
}

func NewMysqlConnection(conf *config.MysqlConfig) *sql.DB {
	mysqlConfig := mysql.NewConfig()
	location, err := time.LoadLocation(Location)
	if err != nil {
		log.Panic("load location error", "error:", err.Error())
		return nil
	}
	{
		mysqlConfig.ParseTime = true
		mysqlConfig.ClientFoundRows = true
		mysqlConfig.InterpolateParams = true
		mysqlConfig.Net = "tcp"
		mysqlConfig.Loc = location
		mysqlConfig.Timeout = time.Second * 5
		mysqlConfig.Collation = "utf8mb4_bin"
	}
	mysqlConfig.User = conf.User
	mysqlConfig.Passwd = conf.Password
	mysqlConfig.DBName = conf.DBName
	mysqlConfig.Addr = conf.Addr
	dsn := mysqlConfig.FormatDSN()
	fmt.Println("mysqlDSN", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic("open mysql db error", "error:", err.Error())
		return nil
	}
	db.SetMaxOpenConns(conf.MaxOpen)
	db.SetMaxIdleConns(conf.MaxIdle)
	db.SetConnMaxLifetime(time.Minute)
	if err := db.Ping(); err != nil {
		log.Panic("mysql ping error", "error:", err.Error())
		return nil
	}
	return db
}