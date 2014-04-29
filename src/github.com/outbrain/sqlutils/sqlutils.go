package sqlutils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/outbrain/config"
	"github.com/outbrain/log"	
)

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

func ScanRowsToMaps(rows *sql.Rows, on_row func(map[string]string) error) error {
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


func QueryRowsMap(db *sql.DB, query string, on_row func(map[string]string) error) error {
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

