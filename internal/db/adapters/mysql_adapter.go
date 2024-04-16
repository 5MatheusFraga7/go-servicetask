package adapters

import (
	"database/sql"
	"fmt"
	"task-manager/internal/db"
)

type MySQLAdapter struct {
	db *sql.DB
}

func NewMySQLAdapter() *MySQLAdapter {
	return &MySQLAdapter{}
}

func (m *MySQLAdapter) Connect() error {

	db, err := sql.Open("mysql", GetConnectionString())
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco de dados MySQL: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("erro ao testar a conexão com o banco de dados MySQL: %v", err)
	}

	m.db = db

	return nil
}

func (m *MySQLAdapter) Close() error {
	if m.db != nil {
		return m.db.Close()
	}
	return nil
}

var _ db.Database = (*MySQLAdapter)(nil)
