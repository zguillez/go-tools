package database

import (
	"database/sql"
	"fmt"

	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"

	"zguillez/go-tools/system"
)

const ECHOSQL = "[sql] %v"

func MySQL(host, user, password, database string, verbose bool) *sql.DB {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, database)
	db, err := sql.Open("mysql", dataSourceName)
	system.CheckError(err)
	if verbose {
		color.Green("*** database connected ***")
		color.Green(fmt.Sprintf("*** [%s:%s]", host, database))
	}
	return db
}

func Close(db *sql.DB, verbose bool) {
	db.Close()
	if verbose {
		color.Yellow("*** database closed ***")
	}
}

func Ping(db *sql.DB, verbose bool) error {
	err := db.Ping()
	if verbose {
		color.Yellow("*** database ping ***")
	}
	return err
}

func Query(db *sql.DB, sql string, verbose bool) (*sql.Rows, []string) {
	return queryHandler(db, sql, verbose)
}
func Queryf(db *sql.DB, sql string, args []string, verbose bool) (*sql.Rows, []string) {
	argx := make([]interface{}, len(args))
	for i, v := range args {
		argx[i] = v
	}
	sql = fmt.Sprintf(fmt.Sprintf(sql, argx...))
	return queryHandler(db, sql, verbose)
}
func queryHandler(db *sql.DB, sql string, verbose bool) (*sql.Rows, []string) {
	if verbose {
		color.Yellow(ECHOSQL, sql)
	}
	rows, err := db.Query(sql)
	system.CheckError(err)

	cols, err := rows.Columns()
	system.CheckError(err)
	return rows, cols
}

func Select(db *sql.DB, sql string, verbose bool) []map[string]string {
	rows, cols := Query(db, sql, verbose)
	return selectHandler(rows, cols)
}
func Selectf(db *sql.DB, sql string, args []string, verbose bool) []map[string]string {
	rows, cols := Queryf(db, sql, args, verbose)
	return selectHandler(rows, cols)
}
func selectHandler(rows *sql.Rows, cols []string) []map[string]string {
	var data []map[string]string
	for rows.Next() {
		item := make(map[string]string)
		columns := make([]string, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		rows.Scan(columnPointers...)
		for i, colName := range cols {
			item[colName] = columns[i]
		}
		data = append(data, item)
	}
	return data
}

func Insert(db *sql.DB, sql string, verbose bool) int {
	rows, err := db.Exec(sql)
	if verbose {
		color.Yellow(ECHOSQL, sql)
	}
	system.CheckError(err)

	id, err := rows.LastInsertId()
	system.CheckError(err)

	return int(id)
}

func Insertf(db *sql.DB, sql string, args []string, verbose bool) int {
	argx := make([]interface{}, len(args))
	for i, v := range args {
		argx[i] = v
	}
	sql = fmt.Sprintf(fmt.Sprintf(sql, argx...))
	rows, err := db.Exec(sql)
	if verbose {
		color.Yellow(ECHOSQL, sql)
	}
	system.CheckError(err)

	id, err := rows.LastInsertId()
	system.CheckError(err)

	return int(id)
}
