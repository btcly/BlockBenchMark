package mysql

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
)

// 数据库的实例信息
type MySQLendpoint struct {
	UserName  string
	Password  string
	IpAddrees string
	Port      int
	DbName    string
	Charset   string
}

// 数据表的结构体
type TableInfo struct {
	Tran_hex     string
	Server_uuid  string
	Client_uuid  string
	Block_name   string
	ChaincodeID  string
	Create_time  int64
	Start_time   int64
	End_time     int64
	Valid        int64
	Trans_count  int64
	Block_height int64
}

var Mysql_endpoint MySQLendpoint

type DBConn struct {
	Db *sqlx.DB
}

var (
	Dbconn            DBConn
	insert_block_name = "block_trans_info"
)

// var table_name string

// 添加不存在就创建数据表
func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/mysql?charset=%s&loc=Local", Mysql_endpoint.UserName, Mysql_endpoint.Password, Mysql_endpoint.IpAddrees, Mysql_endpoint.Port, Mysql_endpoint.Charset)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		glog.Info(fmt.Sprintf("mysql connect failed, detail is [%v]", err.Error()))
	}
	if err := Db.Ping(); err != nil {
		glog.Info(fmt.Sprintf("unable to reach database: %v", err))
	}

	createDBSQL := "CREATE DATABASE IF NOT EXISTS " + Mysql_endpoint.DbName
	_, err = Db.Exec(createDBSQL)
	if err != nil {
		glog.Info(fmt.Sprintf("CREATE database err, %v", err))
	}

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=Local", Mysql_endpoint.UserName, Mysql_endpoint.Password, Mysql_endpoint.IpAddrees, Mysql_endpoint.Port, Mysql_endpoint.DbName, Mysql_endpoint.Charset)
	Db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		glog.Info(fmt.Sprintf("mysql connect failed, detail is [%v]", err.Error()))
	}
	useDBSQL := "USE " + Mysql_endpoint.DbName
	_, err = Db.Exec(useDBSQL)
	if err != nil {
		glog.Exit("use database err, %v", err)
	}

	// Maximum Idle Connections
	Db.SetMaxIdleConns(5)
	// Maximum Open Connections
	Db.SetMaxOpenConns(10)

	// Connection Lifetime
	Db.SetConnMaxLifetime(60 * time.Second)
	// Idle Connection Timeout
	Db.SetConnMaxIdleTime(1 * time.Second)

	//  初始化数据表
	// 如果数据表不在就创建
	create_table_sql(Db)
	Dbconn.Db = Db

	glog.Info(fmt.Sprintf("init mysql Success."))
}

func (dbconn *DBConn) CloseMysql() {
	defer dbconn.Db.Close()
}

//	初始化数据表
//
// 如果数据表不在就创建
func create_table_sql(Db *sqlx.DB) {

	create_sql := "CREATE TABLE IF NOT EXISTS `" + insert_block_name + "` ( " +
		"`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT," +
		"`trans_hex` varchar(200) NOT NULL," +
		"`client_uuid` varchar(200) NOT NULL," +
		"`server_uuid` varchar(200) NOT NULL," +
		"`block_name` varchar(200) NOT NULL," +
		"`chaincodeID` varchar(200) NOT NULL," +
		"`create_time` bigint(20) NOT NULL," +
		"`start_time` bigint(20) NOT NULL," +
		"`end_time` bigint(20) NOT NULL," +
		"`valid` int(11) NOT NULL," +
		"`trans_count` bigint(20) NOT NULL," +
		"`block_height` bigint(20) NOT NULL," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB AUTO_INCREMENT=940892 DEFAULT CHARSET=utf8mb4"

	// glog.Info(fmt.Sprintf("---->result sql:[%s].\n", result_sql)
	_, err := Db.Exec(create_sql)
	if err != nil {
		glog.Info(fmt.Sprintf("mysql CreateBlockTables failed, sql is [%v], err:%s", create_sql, err))
	}
}

// 数据库的字段
func (dbconn *DBConn) InsertBatchBlockInfos(data []TableInfo) {
	sqlStr := "insert into " + insert_block_name + "(trans_hex, server_uuid, client_uuid, block_name, chaincodeID, create_time, start_time, end_time, valid, trans_count, block_height ) values "
	sqlQuery := "(?,?,?,?,?,?,?,?,?,?,?)"

	vals := []interface{}{}
	for index, item := range data {
		if index == len(data)-1 {
			sqlStr += sqlQuery
		} else {
			sqlStr += sqlQuery + ", "
		}

		//  字段赋值
		vals = append(vals,
			item.Tran_hex,
			item.Server_uuid,
			item.Client_uuid,
			item.Block_name,
			item.ChaincodeID,
			item.Create_time,
			item.Start_time,
			item.End_time,
			item.Valid,
			item.Trans_count,
			item.Block_height,
		)

	}

	_, err := dbconn.Db.Exec(sqlStr, vals...) // vals...: 解构
	if err != nil {
		glog.Info(fmt.Sprintf("sqlStr:", sqlStr))
		glog.Info(fmt.Sprintf("vals: ", vals))
		glog.Info(fmt.Sprintf("mysql inert err: ", err))
	}
}
