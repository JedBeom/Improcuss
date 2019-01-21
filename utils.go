package main

import (
	"database/sql"
)

func closeStmt(stmt *sql.Stmt) {
	_ = stmt.Close()
}
