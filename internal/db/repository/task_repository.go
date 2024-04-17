package repository

import (
	"fmt"
	"task-manager/internal/db"
	"task-manager/internal/db/adapters"
	"task-manager/internal/models"
)

type TaskRepository struct {
}

func (tr TaskRepository) CreateTask(t models.Task) error {

	postgresAdapter := adapters.NewPostgreSQLAdapter()
	err := db.OpenConnectionToDatabase(postgresAdapter)

	_, err = postgresAdapter.Exec(getCreateTaskSql(t),
		t.Name,
		t.Description,
		t.UserName,
		t.Deadline,
	)

	defer db.CloseConnectionToDatabase(postgresAdapter)

	if err != nil {
		return fmt.Errorf("falha ao inserir tarefa no banco de dados: %v", err)
	}

	if err != nil {
		panic(err)
	}

	return nil
}
func (tr TaskRepository) readTask() {

}
func (tr TaskRepository) updateTask() {

}
func (tr TaskRepository) deleteTask() {

}

func getCreateTaskSql(t models.Task) string {

	query := `
			INSERT INTO tasks (name, description, username, deadline, done, created_at)
			VALUES ($1, $2, $3, $4, false, now())
	`
	return query
}
