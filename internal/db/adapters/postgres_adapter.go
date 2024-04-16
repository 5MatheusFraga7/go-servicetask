package adapters

import (
	"database/sql"
	"fmt"
	"task-manager/internal/db"

	_ "github.com/lib/pq" // Importa o driver PostgreSQL
)

// PostgreSQLAdapter é a implementação do adaptador para PostgreSQL
type PostgreSQLAdapter struct {
	db *sql.DB
}

// NewPostgreSQLAdapter cria uma nova instância do adaptador PostgreSQL
func NewPostgreSQLAdapter() *PostgreSQLAdapter {
	return &PostgreSQLAdapter{}
}

func (p *PostgreSQLAdapter) Connect() error {

	db, err := sql.Open("postgres", GetConnectionString())
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco de dados PostgreSQL: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("erro ao testar a conexão com o banco de dados PostgreSQL: %v", err)
	}
	p.db = db

	return nil
}

// Close fecha a conexão com o banco de dados PostgreSQL
func (p *PostgreSQLAdapter) Close() error {
	if p.db != nil {
		return p.db.Close()
	}
	return nil
}

// Implemente a interface db.Database no adaptador PostgreSQL
var _ db.Database = (*PostgreSQLAdapter)(nil)
