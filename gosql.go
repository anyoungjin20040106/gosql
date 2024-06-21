package gosql

import (
	"database/sql"
	"fmt"
	"os/exec"
)

type Server struct {
}

func (Server) Mysqlconn(host string, user string, pw string, db string, port int) (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pw, host, port, db))
}

func (Server) Mssqlconn(host string, user string, pw string, db string, port int) (*sql.DB, error) {
	return sql.Open("sqlserver", fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", user, pw, host, port, db))
}

func (Server) Postgresqlconn(host string, user string, pw string, db string, port int) (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pw, db))
}

func (Server) Oracleconn(user string, pw string, host string, port int, db string) (*sql.DB, error) {
	return sql.Open("godror", fmt.Sprintf("%s/%s@%s:%d/%s", user, pw, host, port, db))
}

type Localhost struct {
}

func (Localhost) Mysqlconn(pw string) (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/mysql", pw))
}

func (Localhost) Mssqlconn(pw string) (*sql.DB, error) {
	return sql.Open("sqlserver", fmt.Sprintf("sqlserver://sa:%s@127.0.0.1:1433?database=master", pw))
}

func (Localhost) Postgresqlconn(pw string) (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("host=127.0.0.1 port=5432 user=postgres password=%s dbname=postgres sslmode=disable", pw))
}

func (Localhost) Oracleconn(pw string) (*sql.DB, error) {
	return sql.Open("godror", fmt.Sprintf("system/%s@127.0.0.1:1521/XE", pw))
}

func (Localhost) Sqliteconn(path string) (*sql.DB, error) {
	return sql.Open("sqlite3", path)
}

type Download struct {
}

func (Download) getMysql() bool {
	_, err := exec.Command("go", "get", "github.com/go-sql-driver/mysql").CombinedOutput()
	return err == nil
}
func (Download) getMssql() bool {
	_, err := exec.Command("go", "get", "github.com/denisenkom/go-mssqldb").CombinedOutput()
	return err == nil
}
func (Download) getPostgresql() bool {
	_, err := exec.Command("go", "get", "github.com/lib/pq").CombinedOutput()
	return err == nil
}
func (Download) getOracle() bool {
	_, err := exec.Command("go", "get", "github.com/godror/godror").CombinedOutput()
	return err == nil
}

func (Download) getSqlite() bool {
	_, err := exec.Command("go", "get", "github.com/mattn/go-sqlite3").CombinedOutput()
	return err == nil
}
