package sqlutils

import (
	"fmt"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/outbrain/config"
	"github.com/outbrain/log"	
)


var generateSQL = []string{
	`
		CREATE TABLE database_instance (
		  hostname varchar(128) CHARACTER SET ascii NOT NULL,
		  port smallint(5) unsigned NOT NULL,
		  last_checked timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		  last_seen timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
		  is_last_seen_valid tinyint(3) unsigned NOT NULL,
		  server_id int(10) unsigned NOT NULL,
		  version varchar(128) CHARACTER SET ascii NOT NULL,
		  binlog_format varchar(16) CHARACTER SET ascii NOT NULL,
		  log_bin tinyint(3) unsigned NOT NULL,
		  log_slave_updates tinyint(3) unsigned NOT NULL,
		  binary_log_file varchar(128) CHARACTER SET ascii NOT NULL,
		  binary_log_pos bigint(20) unsigned NOT NULL,
		  master_host varchar(128) CHARACTER SET ascii NOT NULL,
		  master_port smallint(5) unsigned NOT NULL,
		  slave_sql_running tinyint(3) unsigned NOT NULL,
		  slave_io_running tinyint(3) unsigned NOT NULL,
		  master_log_file varchar(128) CHARACTER SET ascii NOT NULL,
		  read_master_log_pos bigint(20) unsigned NOT NULL,
		  relay_master_log_file varchar(128) CHARACTER SET ascii NOT NULL,
		  exec_master_log_pos bigint(20) unsigned NOT NULL,
		  seconds_behind_master bigint(20) unsigned DEFAULT NULL,
		  num_slave_hosts int(10) unsigned NOT NULL,
		  slave_hosts text CHARACTER SET ascii NOT NULL,
		  cluster_name tinytext CHARACTER SET ascii NOT NULL,
		  PRIMARY KEY (hostname,port)
		) ENGINE=InnoDB DEFAULT CHARSET=latin1 ROW_FORMAT=COMPRESSED	
	`,
}

type RowMap map[string]string

func (this *RowMap) GetString(key string) string {
	return (*this)[key]
}

func (this *RowMap) GetInt64(key string) int64 {
	res, _ := strconv.ParseInt((*this)[key], 10, 0)
	return res
}

func (this *RowMap) GetInt(key string) int {
	res, _ := strconv.Atoi((*this)[key])
	return res
}

func (this *RowMap) GetIntD(key string, def int) int {
	res, err := strconv.Atoi((*this)[key])
	if err != nil {return def}
	return res
}

func (this *RowMap) GetUint(key string) uint {
	res, _ := strconv.Atoi((*this)[key])
	return uint(res)
}

func (this *RowMap) GetBool(key string) bool {
	return (*this)[key] == "1"
}


var knownDBs map[string]*sql.DB = make(map[string]*sql.DB)

func getDB(mysql_uri string) (*sql.DB, error) {
	
	if _, exists := knownDBs[mysql_uri]; !exists {
	    if db, err := sql.Open("mysql", mysql_uri); err == nil {
	    	knownDBs[mysql_uri] = db
			log.Info("sqlutils.getDB", mysql_uri)
	    } else {
	    	return db, err
	    }	    	    
	}
	return knownDBs[mysql_uri], nil
}


func OpenTopology(host string, port int) (*sql.DB, error) {
	mysql_uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/", config.Config.MySQLTopologyUser, config.Config.MySQLTopologyPassword, host, port)
	return getDB(mysql_uri)
}

func OpenOrchestrator() (*sql.DB, error) {
	mysql_uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Config.MySQLOrchestratorUser, config.Config.MySQLOrchestratorPassword, 
		config.Config.MySQLOrchestratorHost, config.Config.MySQLOrchestratorPort, config.Config.MySQLOrchestratorDatabase)
	return getDB(mysql_uri)
}

func ExecOrchestrator(query string, args ...interface{}) (sql.Result, error) {
	db,	err	:=	OpenOrchestrator()
	if err != nil {
		return nil, err
	}
	res, err := Exec(db, query, args...)
	return res, err
}

func RowToArray(rows *sql.Rows, columns []string) []string {
    buff := make([]interface{}, len(columns))
    data := make([]string, len(columns))
    for i, _ := range buff {
        buff[i] = &data[i]
    }
	rows.Scan(buff...)
	return data
}

func ScanRowsToArrays(rows *sql.Rows, on_row func([]string) error) error {
    columns, _ := rows.Columns()
    for rows.Next() {
    	arr := RowToArray(rows, columns)
	    
	    err := on_row(arr)
	    if err != nil {
	    	return err
	    }
    }
    return nil
}

func ScanRowsToMaps(rows *sql.Rows, on_row func(RowMap) error) error {
	columns, _ := rows.Columns()
	err := ScanRowsToArrays(rows, func(arr []string) error {
    	m := make(map[string]string)	 
	    for k, data_col := range arr {
	        m[columns[k]] = data_col
	    }
	    err := on_row(m)
	    if err != nil {
	    	return err
	    }
	    return nil
	})
	return err
}


func QueryRowsMap(db *sql.DB, query string, on_row func(RowMap) error) error {
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil && err != sql.ErrNoRows {
		return log.Errore(err)
	} 	
	err = ScanRowsToMaps(rows, on_row)
	return err
}



func Exec(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
    stmt, err := db.Prepare(query);
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var res sql.Result
	res, err = stmt.Exec(args...)
	if err != nil {
		log.Errore(err)
	}
	return res, err
}

