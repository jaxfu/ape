package internal

import (
	"database/sql"
	"os"
)

func initDb(dbPath, initSqlPath string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	if err := runInitDbSql(conn, initSqlPath); err != nil {
		return nil, err
	}

	return conn, nil
}

func runInitDbSql(conn *sql.DB, s string) error {
	sql, err := os.ReadFile(s)
	if err != nil {
		return err
	}

	if _, err := conn.Exec(string(sql)); err != nil {
		return err
	}

	return nil
}
