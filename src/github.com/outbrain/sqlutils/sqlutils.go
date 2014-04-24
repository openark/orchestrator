package sqlutils

import (
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/outbrain/config"
)

func OpenTopology(host string, port int) (*sql.DB, error) {
	mysql_uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/", config.Config.MySQLTopologyUser, config.Config.MySQLTopologyPassword, host, port)
	log.Println("sqlutils.Open", mysql_uri)
	return sql.Open("mysql", mysql_uri)
}

func OpenOrchestrator() (*sql.DB, error) {
	mysql_uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Config.MySQLOrchestratorUser, config.Config.MySQLOrchestratorPassword, 
		config.Config.MySQLOrchestratorHost, config.Config.MySQLOrchestratorPort, config.Config.MySQLOrchestratorDatabase)
	log.Println("sqlutils.Open", mysql_uri)
	return sql.Open("mysql", mysql_uri)
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

func ScanRowsToArrays(rows *sql.Rows, on_row func([]string)) {
    columns, _ := rows.Columns()
    for rows.Next() {
    	arr := RowToArray(rows, columns)
	    on_row(arr)
    }
}

func ScanRowsToMaps(rows *sql.Rows, on_row func(map[string]string)) {
	columns, _ := rows.Columns()
	ScanRowsToArrays(rows, func(arr []string) {
    	m := make(map[string]string)	 
	    for k, data_col := range arr {
	        m[columns[k]] = data_col
	    }
	    on_row(m)
	})
}


func QueryRowsMap(db *sql.DB, query string,  on_row func(map[string]string)) error {
	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return err
	} 	
	ScanRowsToMaps(rows, on_row)
	return nil
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
		log.Println(err)
	}
	return res, err
}

